package user_repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"smartway-test/database/user_repo/query"
	"smartway-test/model"
	"strconv"
	"strings"
	"time"
)

type UserRepository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewUserRepository(logger *zap.SugaredLogger, db *sqlx.DB) UserRepository {
	return UserRepository{
		logger: logger,
		db:     db,
	}
}

func (r UserRepository) InsertUser(ctx context.Context, userModel model.User) error {
	_, err := r.db.ExecContext(ctx, query.InsertUserSQL, userModel.Name, userModel.Surname, userModel.Patronymic)
	if err != nil {
		return errors.New(fmt.Sprintf("error on insert user: %s", err.Error()))
	}

	return nil
}

func (r UserRepository) CheckUser(ctx context.Context, userId int) (bool, error) {
	var count int

	err := r.db.QueryRowContext(ctx, query.CheckUserSql, userId).Scan(&count)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (r UserRepository) UpdateUser(ctx context.Context, userModel model.User, userId string) error {
	q := `update service_user set`
	qParts := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	argId := 1

	if len(userModel.Name) != 0 {
		qParts = append(qParts, fmt.Sprintf(" name = $%d", argId))
		argId++
		args = append(args, userModel.Name)
	}

	if len(userModel.Surname) != 0 {
		qParts = append(qParts, fmt.Sprintf(" surname = $%d", argId))
		argId++
		args = append(args, userModel.Surname)
	}

	if len(userModel.Patronymic) != 0 {
		qParts = append(qParts, fmt.Sprintf(" patronymic = $%d", argId))
		argId++
		args = append(args, userModel.Patronymic)
	}

	q += strings.Join(qParts, ",") + fmt.Sprintf(` where id = $%d`, argId)
	args = append(args, userId)

	_, err := r.db.ExecContext(ctx, q, args...)

	return err
}

func (r UserRepository) DeleteUser(ctx context.Context, userId string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query.DeleteUserTicketsSql)
	if err != nil {
		if maybeErr := tx.Rollback(); maybeErr != nil {
			return maybeErr
		}
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, userId); err != nil {
		if maybeErr := tx.Rollback(); maybeErr != nil {
			return maybeErr
		}
		return err
	}

	stmt, err = tx.PrepareContext(ctx, query.DeleteUserDocumentsSql)
	if err != nil {
		if maybeErr := tx.Rollback(); maybeErr != nil {
			return maybeErr
		}
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, userId); err != nil {
		if maybeErr := tx.Rollback(); maybeErr != nil {
			return maybeErr
		}
		return err
	}

	stmt, err = tx.PrepareContext(ctx, query.DeleteUserSql)
	if err != nil {
		if maybeErr := tx.Rollback(); maybeErr != nil {
			return maybeErr
		}
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, userId); err != nil {
		if maybeErr := tx.Rollback(); maybeErr != nil {
			return maybeErr
		}
		return err
	}

	return tx.Commit()
}

func (r UserRepository) GetInfoAboutUser(ctx context.Context, startTime time.Time, endTime time.Time, userId string) (model.ReportFlights, error) {
	var (
		previouslyTickets             []model.FullTicketInfo
		notFulFilledTickets           []model.FullTicketInfo
		previouslyNotFulFilledTickets []model.FullTicketInfo
	)

	row, err := r.db.QueryContext(ctx, query.SelectPreviouslyFlightsSql, startTime, endTime, userId)
	if err != nil {
		return model.ReportFlights{}, err
	}

	for row.Next() {
		var ticket model.FullTicketInfo

		if err := row.Scan(&ticket.Id, &ticket.StartPoint, &ticket.EndPoint, &ticket.StartTime, &ticket.EndTime, &ticket.BuyTime, &ticket.Company); err != nil {
			return model.ReportFlights{}, err
		}

		userId, _ := strconv.Atoi(userId)
		ticket.UserId = userId

		previouslyTickets = append(previouslyTickets, ticket)
	}

	row, err = r.db.QueryContext(ctx, query.SelectNotFulFilledFlightsSql, startTime, endTime, userId)
	if err != nil {
		return model.ReportFlights{}, err
	}

	for row.Next() {
		var ticket model.FullTicketInfo

		if err := row.Scan(&ticket.Id, &ticket.StartPoint, &ticket.EndPoint, &ticket.StartTime, &ticket.EndTime, &ticket.BuyTime, &ticket.Company); err != nil {
			return model.ReportFlights{}, err
		}

		userId, _ := strconv.Atoi(userId)
		ticket.UserId = userId

		notFulFilledTickets = append(notFulFilledTickets, ticket)
	}

	row, err = r.db.QueryContext(ctx, query.SelectPreviouslyAndNotFulFilledFlightsSql, startTime, endTime, userId)
	if err != nil {
		return model.ReportFlights{}, err
	}

	for row.Next() {
		var ticket model.FullTicketInfo

		if err := row.Scan(&ticket.Id, &ticket.StartPoint, &ticket.EndPoint, &ticket.StartTime, &ticket.EndTime, &ticket.BuyTime, &ticket.Company); err != nil {
			return model.ReportFlights{}, err
		}

		userId, _ := strconv.Atoi(userId)
		ticket.UserId = userId

		previouslyNotFulFilledTickets = append(previouslyNotFulFilledTickets, ticket)
	}

	return model.ReportFlights{
		Previously:                previouslyTickets,
		NotFulFilled:              notFulFilledTickets,
		PreviouslyAndNotFulFilled: previouslyNotFulFilledTickets,
	}, nil
}
