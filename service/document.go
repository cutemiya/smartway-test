package service

import (
	"context"
	"go.uber.org/zap"
	database "smartway-test/database/doc_repo"
	"smartway-test/model"
)

type DocRepo interface {
	InsertDocument(ctx context.Context, userId int, docModel model.Document) error
	SelectAllDocumentsByUserId(ctx context.Context, userId int) ([]model.Document, error)
	UpdateDocument(ctx context.Context, documentModel model.Document, documentId string) error
	DeleteDocument(ctx context.Context, docId string) error
}

type DocService struct {
	logger *zap.SugaredLogger
	repo   database.DocRepository
}

func NewDocService(logger *zap.SugaredLogger, repo database.DocRepository) DocService {
	return DocService{
		logger: logger,
		repo:   repo,
	}
}

func (s DocService) NewDocument(ctx context.Context, userId int, documentModel model.Document) error {
	return s.repo.InsertDocument(ctx, userId, documentModel)
}

func (s DocService) GetAllDocumentByUserId(ctx context.Context, userId int) ([]model.Document, error) {
	return s.repo.SelectAllDocumentsByUserId(ctx, userId)
}

func (s DocService) UpdateDocument(ctx context.Context, documentModel model.Document, documentId string) error {
	return s.repo.UpdateDocument(ctx, documentModel, documentId)
}

func (s DocService) DeleteDocument(ctx context.Context, docId string) error {
	return s.repo.DeleteDocument(ctx, docId)
}
