package user

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/client/user"
	"smartway-test/lib/response"
	"smartway-test/service"
)

// CreateUserHandle
//
//	@Description	create a new document
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		user.User	true	"query params"
//	@Success		200		{object}	client.Success
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/user/new [post]
func CreateUserHandle(logger *zap.SugaredLogger, service service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request user.User

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)

		if err != nil { // 500
			logger.Debugf("Decode request error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: err.Error(),
			})
			response.SendResponse(w, 500, res)
			return
		}

		if len(request.Name) == 0 || len(request.Surname) == 0 || len(request.Patronymic) == 0 { // 400
			logger.Debugf("Fill in the all field, 400")
			res, _ := json.Marshal(client.Error{
				Code:  400,
				Error: "Fill in the all field",
			})
			response.SendResponse(w, 400, res)
			return
		}

		err = service.NewUser(r.Context(), user.MapToUserServiceModel(request))
		if err != nil { // 500
			logger.Debugf("service error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("service error: %s", err.Error()),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("succes /user/new")
		res, _ := json.Marshal(client.Success{
			Ok: true,
		})
		response.SendResponse(w, 200, res)
	}
}
