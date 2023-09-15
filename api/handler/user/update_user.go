package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/client/user"
	"smartway-test/lib/response"
	"smartway-test/service"
)

// UpdateUserHandle
//
//	@Description	update user document
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userId		path		int				true	"user id"
//	@Param			request	body		user.User	true	"query params"
//	@Success		200		{object}	client.Success
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/user/update/{userId}  [patch]
func UpdateUserHandle(logger *zap.SugaredLogger, userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["userId"]
		var request user.User

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

		if err = userService.UpdateUserInfo(r.Context(), user.MapToUserServiceModel(request), userId); err != nil {
			logger.Debugf("Internal serivce error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: err.Error(),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("Success /user/update 200")
		res, _ := json.Marshal(client.Success{Ok: true})
		response.SendResponse(w, 200, res)
	}
}
