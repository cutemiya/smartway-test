package document

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/lib/response"
	"smartway-test/service"
)

// DeleteDocumentHandle
//
//	@Description	delete a document
//	@Tags			docs
//	@Produce		json
//	@Param			userId	path		int	true	"document id"
//	@Success		200	{object}	client.Success
//	@Failure		400	{object}	client.Error
//	@Failure		500	{object}	client.Error
//	@Router			/doc/delete/{docId} [delete]
func DeleteDocumentHandle(logger *zap.SugaredLogger, documentService service.DocService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		docId := mux.Vars(r)["docId"]

		if err := documentService.DeleteDocument(r.Context(), docId); err != nil {
			logger.Debugf("Internal server error: %s", err)
			res, _ := json.Marshal(client.Error{
				Error: fmt.Sprintf("Internal server error: %s", err),
				Code:  500,
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("Success /doc/delete/{docId}")
		res, _ := json.Marshal(client.Success{Ok: true})
		response.SendResponse(w, 200, res)
	}
}
