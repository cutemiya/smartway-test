package doc_repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"smartway-test/database/doc_repo/query"
	"smartway-test/model"
	"strings"
)

type DocRepository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewDocRepository(logger *zap.SugaredLogger, db *sqlx.DB) DocRepository {
	return DocRepository{
		logger: logger,
		db:     db,
	}
}

func (r DocRepository) InsertDocument(ctx context.Context, userId int, docModel model.Document) error {
	_, err := r.db.ExecContext(ctx, query.InsertDocumentSql, docModel.Type, docModel.Number, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r DocRepository) SelectAllDocumentsByUserId(ctx context.Context, userId int) ([]model.Document, error) {
	var docsList []model.Document
	var docType string
	var docNumber string
	var docId int

	row, err := r.db.QueryContext(ctx, query.SelectAllDocumentsByUserId, userId)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&docId, &docType, &docNumber)
		if err != nil {
			return nil, err
		}

		docsList = append(docsList, model.Document{
			Id:     docId,
			Type:   docType,
			Number: docNumber,
		})
	}

	return docsList, nil
}

func (r DocRepository) UpdateDocument(ctx context.Context, documentModel model.Document, documentId string) error {
	q := `update user_document set`
	qParts := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	argId := 1

	if len(documentModel.Type) != 0 {
		qParts = append(qParts, fmt.Sprintf(" doc_type = $%d", argId))
		argId++
		args = append(args, documentModel.Type)
	}

	if len(documentModel.Number) != 0 {
		qParts = append(qParts, fmt.Sprintf(" number = $%d", argId))
		argId++
		args = append(args, documentModel.Number)
	}

	q += strings.Join(qParts, ",") + fmt.Sprintf(` where id = $%d`, argId)
	args = append(args, documentId)

	_, err := r.db.ExecContext(ctx, q, args...)

	return err
}

func (r DocRepository) DeleteDocument(ctx context.Context, docId string) error {
	_, err := r.db.ExecContext(ctx, query.DeleteDocumentSql, docId)
	return err
}
