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

// GetAllUSerDocumentsByUserIdHandle
//
//	@Description	Get All USer Documents By User Id
//	@Tags			docs
//	@Produce		json
//	@Param			userId		path		int				true	"user id"
//	@Success		200		{object}	doc.DocumentList
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/doc/get/user/{userId} [get]
func GetAllUSerDocumentsByUserIdHandle(logger *zap.SugaredLogger, service service.DocService, userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
		isUser, err := userService.HaveUser(r.Context(), userId)
		if err != nil {
			logger.Debugf("internal service error %s, 500", err)
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

		docs, err := service.GetAllDocumentByUserId(r.Context(), userId)
		if err != nil {
			logger.Debugf("internal service error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: err.Error(),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("successful /doc/get/user/{userId}/, 200")
		res, _ := json.Marshal(doc.DocumentList{Documents: doc.MapToClientDocumentModel(docs)})
		response.SendResponse(w, 200, res)
	}
}
