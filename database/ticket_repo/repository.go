package ticket_repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"smartway-test/database/ticket_repo/query"
	"smartway-test/model"
	"strings"
)

type TicketRepository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewDocRepository(logger *zap.SugaredLogger, db *sqlx.DB) TicketRepository {
	return TicketRepository{
		logger: logger,
		db:     db,
	}
}

func (r TicketRepository) InsertTicket(ctx context.Context, userId int, ticketModel model.Ticket) error {
	_, err := r.db.ExecContext(ctx,
		query.InsertTicketSql,
		ticketModel.StartPoint,
		ticketModel.EndPoint,
		//ticketModel.StartDate,
		//ticketModel.EndDate,
		ticketModel.StartTime,
		ticketModel.EndTime,
		ticketModel.Company,
		userId)
	if err != nil {
		return err
	}

	return nil
}

func (r TicketRepository) SelectAllTickets(ctx context.Context) ([]model.FullTicketInfo, error) {
	var (
		tickets []model.FullTicketInfo

		ticketId   int
		startPoint string
		endPoint   string
		startTime  string
		endTime    string
		buyTime    string
		company    string
		userId     int
	)
	row, err := r.db.QueryContext(ctx, query.SelectAllTickets)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&ticketId, &startPoint, &endPoint, &startTime, &endTime, &buyTime, &company, &userId)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, model.FullTicketInfo{
			Id:      ticketId,
			BuyTime: buyTime,
			UserId:  userId,
			Ticket: model.Ticket{
				StartPoint: startPoint,
				EndPoint:   endPoint,
				StartTime:  startTime,
				EndTime:    endTime,
				Company:    company,
			},
		})
	}

	return tickets, nil
}

func (r TicketRepository) SelectAllPassengerByTicketId(ctx context.Context, ticketId int) ([]model.User, error) {
	var (
		passengerList []model.User

		name       string
		surname    string
		patronymic string
	)

	row, err := r.db.QueryContext(ctx, query.SelectUsersByTicketId, ticketId)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(&name, &surname, &patronymic)
		if err != nil {
			return nil, err
		}

		passengerList = append(passengerList, model.User{
			Name:       name,
			Surname:    surname,
			Patronymic: patronymic,
		})
	}

	return passengerList, nil
}

func (r TicketRepository) SelectAllInfoByTicket(ctx context.Context, ticketId string) (model.AllInfoOfTicket, error) {
	var (
		out        model.AllInfoOfTicket
		userId     int
		name       string
		surname    string
		patronymic string

		passengers []model.Passenger
	)

	if err := r.db.QueryRowContext(ctx, query.SelectAllInfoAboutTicketByIdSql, ticketId).
		Scan(&out.StartPoint, &out.EndPoint,
			&out.StartTime, &out.EndTime, &out.BuyTime,
			&out.Company); err != nil {
		return model.AllInfoOfTicket{}, err
	}

	row, err := r.db.QueryContext(ctx, query.SelectAllInfoAboutUserSql, ticketId)
	if err != nil {
		return model.AllInfoOfTicket{}, err
	}

	for row.Next() {
		if err = row.Scan(&userId, &name, &surname, &patronymic); err != nil {
			return model.AllInfoOfTicket{}, err
		}

		var documents []model.Document
		var docNumber string
		var docType string
		var docId int

		docRow, err := r.db.QueryContext(ctx, query.SelectAllDocumentsByUserId, userId)
		if err != nil {
			return model.AllInfoOfTicket{}, err
		}

		for docRow.Next() {
			if err = docRow.Scan(&docId, &docType, &docNumber); err != nil {
				return model.AllInfoOfTicket{}, err
			}

			documents = append(documents, model.Document{
				Id:     docId,
				Type:   docType,
				Number: docNumber,
			})
		}

		var newPassenger = model.Passenger{
			User: model.User{
				Name:       name,
				Surname:    surname,
				Patronymic: patronymic,
			},
			Documents: documents,
		}

		passengers = append(passengers, newPassenger)
	}

	out.Passengers = passengers

	return out, nil
}

// update

func (r TicketRepository) UpdateTicket(ctx context.Context, ticketModel model.FullTicketInfo, ticketId string) error {
	q := `update tickets set`
	qParts := make([]string, 0, 10)
	args := make([]interface{}, 0, 10)
	argId := 1

	if len(ticketModel.StartPoint) != 0 {
		qParts = append(qParts, fmt.Sprintf(" start_point = $%d", argId))
		argId++
		args = append(args, ticketModel.StartPoint)
	}

	if len(ticketModel.EndPoint) != 0 {
		qParts = append(qParts, fmt.Sprintf(" end_point = $%d", argId))
		argId++
		args = append(args, ticketModel.EndPoint)
	}

	if len(ticketModel.StartTime) != 0 {
		qParts = append(qParts, fmt.Sprintf(" start_time = $%d", argId))
		argId++
		args = append(args, ticketModel.StartTime)
	}

	if len(ticketModel.EndTime) != 0 {
		qParts = append(qParts, fmt.Sprintf(" end_time = $%d", argId))
		argId++
		args = append(args, ticketModel.EndTime)
	}

	if len(ticketModel.BuyTime) != 0 {
		qParts = append(qParts, fmt.Sprintf(" buy_time = $%d", argId))
		argId++
		args = append(args, ticketModel.BuyTime)
	}

	if len(ticketModel.Company) != 0 {
		qParts = append(qParts, fmt.Sprintf(" company = $%d", argId))
		argId++
		args = append(args, ticketModel.Company)
	}

	if ticketModel.UserId > 0 {
		qParts = append(qParts, fmt.Sprintf(" user_id = $%d", argId))
		argId++
		args = append(args, ticketModel.UserId)
	}

	q += strings.Join(qParts, ",") + fmt.Sprintf(` where id = $%d`, argId)
	args = append(args, ticketId)

	_, err := r.db.ExecContext(ctx, q, args...)

	return err
}

// delete

func (r TicketRepository) DeleteTicket(ctx context.Context, ticketId string) error {
	_, err := r.db.ExecContext(ctx, query.DeleteTicketSql, ticketId)
	return err
}
