package document

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/client/doc"
	"smartway-test/lib/response"
	"smartway-test/service"
)

// UpdateDocumentHandle
//
//	@Description	update user document
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			docId		path		int				true	"document id"
//	@Param			request	body		doc.Document	true	"query params"
//	@Success		200		{object}	client.Success
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/doc/update/{docId} [patch]
func UpdateDocumentHandle(logger *zap.SugaredLogger, documentService service.DocService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		docId := mux.Vars(r)["docId"]
		var request doc.Document

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: err.Error(),
			})
			response.SendResponse(w, 500, res)
			return
		}

		if err = documentService.UpdateDocument(r.Context(), doc.MapToDocumentServiceModel(request), docId); err != nil {
			logger.Debugf("Internal serivce error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: err.Error(),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("Success /doc/update: %s", err)
		res, _ := json.Marshal(client.Success{Ok: true})
		response.SendResponse(w, 200, res)
	}
}
