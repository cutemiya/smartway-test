package ticket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/client/ticket"
	"smartway-test/lib/response"
	"smartway-test/service"
	"strconv"
)

// NewTicketHandle
//
//	@Description	create a new ticket
//	@Tags			ticket
//	@Accept			json
//	@Produce		json
//	@Param			userId		path		int				true	"user id"
//	@Param			request	body		ticket.Ticket	true	"query params"
//	@Success		200		{object}	client.Success
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/ticket/new/{userId} [post]
func NewTicketHandle(logger *zap.SugaredLogger, ticketService service.TicketService, userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request ticket.Ticket
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

		if len(request.StartPoint) == 0 || len(request.EndPoint) == 0 || // len(request.StartDate) == 0 ||
			//len(request.EndDate) == 0 ||
			len(request.Company) == 0 || len(request.StartTime) == 0 ||
			len(request.EndTime) == 0 { // 400
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

		err = ticketService.CreateTicket(r.Context(), userId, ticket.MapToServiceTicketModel(request))
		if err != nil {
			logger.Debugf("internal service error %s, 500", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("internal service error: %s", err),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("success /ticket/new/{userId}, 200")
		res, _ := json.Marshal(client.Success{
			Ok: true,
		})
		response.SendResponse(w, 200, res)
	}
}
