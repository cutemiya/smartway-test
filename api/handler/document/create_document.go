package document

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/client/doc"
	"smartway-test/lib/response"
	"smartway-test/service"
	"strconv"
)

// NewDocumentHandle
//
//	@Description	create a new document
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			userId		path		int				true	"user id"
//	@Param			request	body		doc.Document	true	"query params"
//	@Success		200		{object}	client.Success
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/doc/new/{userId} [post]
func NewDocumentHandle(logger *zap.SugaredLogger, docService service.DocService, userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request doc.Document
		userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

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

		if len(request.Type) == 0 || len(request.Number) == 0 { // 400
			logger.Debugf("Fill in the all field, 400")
			res, _ := json.Marshal(client.Error{
				Code:  400,
				Error: "Fill in the all field",
			})
			response.SendResponse(w, 400, res)
			return
		}

		isUser, err := userService.HaveUser(r.Context(), userId)
		if err != nil {
			logger.Debugf("internal service error %s, 400", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("internal service error %s", err),
			})
			response.SendResponse(w, 500, res)
			return
		}

		if !isUser {
			logger.Debugf("user not found, 400")
			res, _ := json.Marshal(client.Error{
				Code:  400,
				Error: "user not found",
			})
			response.SendResponse(w, 400, res)
			return
		}

		err = docService.NewDocument(r.Context(), userId, doc.MapToDocumentServiceModel(request))
		if err != nil {
			logger.Debugf("internal service error, 400")
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("internal service error %s", err),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("successful /doc/new/{userId}, 200")
		res, _ := json.Marshal(client.Success{
			Ok: true,
		})
		response.SendResponse(w, 200, res)
	}
}
