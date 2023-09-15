package service

import (
	"context"
	"go.uber.org/zap"
	database "smartway-test/database/user_repo"
	"smartway-test/model"
	"time"
)

type UserRepo interface {
	InsertUser(ctx context.Context, model model.User) error
	CheckUser(ctx context.Context, userId int) (bool, error)
	UpdateUser(ctx context.Context, userModel model.User, userId string) error
	DeleteUser(ctx context.Context, userId string) error
	GetInfoAboutUser(ctx context.Context, startTime time.Time, endTime time.Time, userId string) (model.ReportFlights, error)
}

type UserService struct {
	logger *zap.SugaredLogger
	repo   database.UserRepository
}

func NewUserService(logger *zap.SugaredLogger, repo database.UserRepository) UserService {
	return UserService{
		logger: logger,
		repo:   repo,
	}
}

func (s UserService) NewUser(ctx context.Context, user model.User) error {
	return s.repo.InsertUser(ctx, user)
}

func (s UserService) HaveUser(ctx context.Context, userId int) (bool, error) {
	return s.repo.CheckUser(ctx, userId)
}

func (s UserService) UpdateUserInfo(ctx context.Context, userModel model.User, userId string) error {
	return s.repo.UpdateUser(ctx, userModel, userId)
}

func (s UserService) DeleteUser(ctx context.Context, userId string) error {
	return s.repo.DeleteUser(ctx, userId)
}

func (s UserService) GetAllInfoOnTimeDiapason(ctx context.Context, startTime time.Time, endTime time.Time, userId string) (model.ReportFlights, error) {
	return s.repo.GetInfoAboutUser(ctx, startTime, endTime, userId)
}
