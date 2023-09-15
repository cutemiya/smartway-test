package service

import (
	"context"
	"go.uber.org/zap"
	database "smartway-test/database/ticket_repo"
	"smartway-test/model"
)

type TicketRepo interface {
	InsertTicket(ctx context.Context, userId int, ticketModel model.Ticket) error
	SelectAllTickets(ctx context.Context) ([]model.FullTicketInfo, error)
	DeleteTicket(ctx context.Context, ticketId string) error
	UpdateTicket(ctx context.Context, ticketModel model.FullTicketInfo, ticketId string) error
	SelectAllPassengerByTicketId(ctx context.Context, ticketId int) ([]model.User, error)
	SelectAllInfoByTicket(ctx context.Context, ticketId string) (model.AllInfoOfTicket, error)
}

type TicketService struct {
	logger *zap.SugaredLogger
	repo   database.TicketRepository
}

func NewTicketService(logger *zap.SugaredLogger, repo database.TicketRepository) TicketService {
	return TicketService{
		logger: logger,
		repo:   repo,
	}
}

func (s TicketService) CreateTicket(ctx context.Context, userId int, ticketModel model.Ticket) error {
	return s.repo.InsertTicket(ctx, userId, ticketModel)
}

func (s TicketService) GetAllTickets(ctx context.Context) ([]model.FullTicketInfo, error) {
	return s.repo.SelectAllTickets(ctx)
}

func (s TicketService) GetAllPassengersByTicketId(ctx context.Context, ticketId int) ([]model.User, error) {
	return s.repo.SelectAllPassengerByTicketId(ctx, ticketId)
}

func (s TicketService) UpdateTicket(ctx context.Context, ticketModel model.FullTicketInfo, ticketId string) error {
	return s.repo.UpdateTicket(ctx, ticketModel, ticketId)
}

func (s TicketService) DeleteTicket(ctx context.Context, ticketId string) error {
	return s.repo.DeleteTicket(ctx, ticketId)
}

func (s TicketService) GetAllInfoByTicketId(ctx context.Context, ticketId string) (model.AllInfoOfTicket, error) {
	return s.repo.SelectAllInfoByTicket(ctx, ticketId)
}
