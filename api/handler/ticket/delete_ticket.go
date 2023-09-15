package ticket

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

// DeleteTicketHandle
//
//	@Description	delete a ticket
//	@Tags			ticket
//	@Produce		json
//	@Param			ticketId	path		int	true	"ticket id"
//	@Success		200	{object}	client.Success
//	@Failure		400	{object}	client.Error
//	@Failure		500	{object}	client.Error
//	@Router			/ticket/delete/{ticketId} [delete]
func DeleteTicketHandle(logger *zap.SugaredLogger, ticketService service.TicketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ticketId := mux.Vars(r)["ticketId"]
		if err := ticketService.DeleteTicket(r.Context(), ticketId); err != nil {
			logger.Debugf("Internal server error: %s", err)
			res, _ := json.Marshal(client.Error{
				Error: fmt.Sprintf("Internal server error: %s", err),
				Code:  500,
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("Success /ticket/delete/{ticketId}")
		res, _ := json.Marshal(client.Success{Ok: true})
		response.SendResponse(w, 200, res)
	}
}
