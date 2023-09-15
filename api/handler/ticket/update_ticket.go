package ticket

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"smartway-test/client"
	"smartway-test/client/ticket"
	"smartway-test/lib/response"
	"smartway-test/service"
)

// UpdateTicketHandle
//
//	@Description	update a ticket
//	@Tags			ticket
//	@Accept			json
//	@Produce		json
//	@Param			ticketId		path		int				true	"ticket id"
//	@Param			request	body		ticket.FullInfoTicket	true	"query params"
//	@Success		200		{object}	client.Success
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/ticket/update/{ticketId} [patch]
func UpdateTicketHandle(logger *zap.SugaredLogger, ticketService service.TicketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		docId := mux.Vars(r)["ticketId"]
		var request ticket.FullInfoTicket

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

		if err = ticketService.UpdateTicket(r.Context(), ticket.MapToServiceFullTicket(request), docId); err != nil {
			logger.Debugf("Internal serivce error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: err.Error(),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("Success /ticket/update: %s", err)
		res, _ := json.Marshal(client.Success{Ok: true})
		response.SendResponse(w, 200, res)
	}
}
