package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/client/user"
	"smartway-test/lib/response"
	"smartway-test/service"
	"time"
)

// GetInfoAboutUsersHandle
//
//		@Description	Get a report of user for diapason
//		@Tags			user
//	 @Accept			json
//		@Produce		json
//		@Param			userId		path		int				true	"user id"
//		@Param			request	body		user.TimeDiapason	true	"query params"
//		@Success		200		{object}	model.ReportFlights
//		@Failure		400		{object}	client.Error
//		@Failure		500		{object}	client.Error
//		@Router			/user/get/{userId} [post]
func GetInfoAboutUsersHandle(logger *zap.SugaredLogger, userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["userId"]

		var req user.TimeDiapason
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: err.Error(),
			})
			response.SendResponse(w, 500, res)
			return
		}
		//RFC3339
		layout := "2006-01-02T15:04:05Z07:00"

		startTime, err := time.Parse(layout, req.StartTime)
		if err != nil {
			res, _ := json.Marshal(client.Error{
				Code:  400,
				Error: "the time should be in the format 2006-01-02T15:04:05Z, where 2006 - year, 01 - month, 02 - day, 15 - hour, 04 - minutes, 05 - seconds",
			})
			response.SendResponse(w, 400, res)
			return
		}

		endTime, err := time.Parse(layout, req.EndTime)
		if err != nil {
			logger.Debugf("the time should be in the format 2006-01-02T15:04:05Z, where 2006 - year, 01 - month, 02 - day, 15 - hour, 04 - minutes, 05 - seconds")
			res, _ := json.Marshal(client.Error{
				Code:  400,
				Error: "the time should be in the format 2006-01-02T15:04:05Z, where 2006 - year, 01 - month, 02 - day, 15 - hour, 04 - minutes, 05 - seconds",
			})
			response.SendResponse(w, 400, res)
			return
		}

		tickets, err := userService.GetAllInfoOnTimeDiapason(r.Context(), startTime, endTime, userId)
		if err != nil {
			logger.Debugf("internal server error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("internal server error: %s", err),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("success user/get/{userId} 200")
		res, _ := json.Marshal(tickets)
		response.SendResponse(w, 200, res)
	}
}
