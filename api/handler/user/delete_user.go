package user

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

// DeleteUserHandle
//
//	@Description	delete a user
//	@Tags			user
//	@Produce		json
//	@Param			userId	path		int	true	"user id"
//	@Success		200	{object}	client.Success
//	@Failure		400	{object}	client.Error
//	@Failure		500	{object}	client.Error
//	@Router			/user/delete/{userId} [delete]
func DeleteUserHandle(logger *zap.SugaredLogger, userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["userId"]

		if err := userService.DeleteUser(r.Context(), userId); err != nil {
			logger.Debugf("Internal server error: %s", err)
			res, _ := json.Marshal(client.Error{
				Error: fmt.Sprintf("Internal server error: %s", err),
				Code:  500,
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("Success /user/delete/{userId}")
		res, _ := json.Marshal(client.Success{Ok: true})
		response.SendResponse(w, 200, res)
	}
}
