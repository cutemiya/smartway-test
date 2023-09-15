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

// GetAllTicketsHandle
//
//	@Description	get all tickets
//	@Tags			ticket
//	@Produce		json
//	@Success		200		{object}	ticket.FullInfoTicketList
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/ticket/all [get]
func GetAllTicketsHandle(logger *zap.SugaredLogger, ticketService service.TicketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickets, err := ticketService.GetAllTickets(r.Context())
		if err != nil {
			logger.Debugf("internal service error: %s, 500", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("internal service error: %s", err),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("success /tickets/all 200")
		res, _ := json.Marshal(ticket.MapToClientTicketList(tickets))
		response.SendResponse(w, 200, res)
	}
}

// GetAllPassengersByTicketIdHandle
//
//	@Description	get all passengers by ticket id
//	@Tags			ticket
//	@Produce		json
//	@Param			ticketId		path		int				true	"ticket id"
//	@Success		200		{object}	ticket.PassengersByTicketId
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/ticket/passengers/{ticketId} [get]
func GetAllPassengersByTicketIdHandle(logger *zap.SugaredLogger, ticketService service.TicketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ticketId, _ := strconv.Atoi(mux.Vars(r)["ticketId"])
		passengers, err := ticketService.GetAllPassengersByTicketId(r.Context(), ticketId)
		if err != nil {
			logger.Debugf("internal service error %s, 500", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("internal service error %s", err),
			})
			response.SendResponse(w, 500, res)
			return
		}
		logger.Debugf("success /tickets/passengers/{ticketId}, 200")
		res, _ := json.Marshal(ticket.MapToClientPassengers(passengers))
		response.SendResponse(w, 200, res)
	}
}

// GetAllInfoFromTicketIdHandle
//
//	@Description	get all info about ticket from ticket id
//	@Tags			ticket
//	@Produce		json
//	@Param			ticketId		path		int				true	"ticket id"
//	@Success		200		{object}	ticket.AllINfoAboutTicket
//	@Failure		400		{object}	client.Error
//	@Failure		500		{object}	client.Error
//	@Router			/tickets/get/all/{ticketId} [get]
func GetAllInfoFromTicketIdHandle(logger *zap.SugaredLogger, ticketService service.TicketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ticketId := mux.Vars(r)["ticketId"]

		allInfo, err := ticketService.GetAllInfoByTicketId(r.Context(), ticketId)
		if err != nil {
			logger.Debugf("Internal server error: %s", err)
			res, _ := json.Marshal(client.Error{
				Code:  500,
				Error: fmt.Sprintf("Internal server error: %s", err),
			})
			response.SendResponse(w, 500, res)
			return
		}

		logger.Debugf("success /tickets/get/all/{ticketId}, 200")
		res, _ := json.Marshal(ticket.MapToClientAllInfoAboutTicket(allInfo))
		response.SendResponse(w, 200, res)
	}
}
