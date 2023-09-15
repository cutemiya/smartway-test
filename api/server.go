package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"smartway-test/api/handler"
	"smartway-test/api/handler/document"
	"smartway-test/api/handler/ticket"
	"smartway-test/api/handler/user"
	"smartway-test/config"
	"smartway-test/lib/pctx"
	"smartway-test/service"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	userService service.UserService,
	docService service.DocService,
	ticketService service.TicketService,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.Ping()).Methods(http.MethodGet)

	router.HandleFunc("/user/new", user.CreateUserHandle(logger, userService)).Methods(http.MethodPost)
	router.HandleFunc("/user/update/{userId}", user.UpdateUserHandle(logger, userService)).Methods(http.MethodPatch)
	router.HandleFunc("/user/delete/{userId}", user.DeleteUserHandle(logger, userService)).Methods(http.MethodDelete)
	router.HandleFunc("/user/get/{userId}", user.GetInfoAboutUsersHandle(logger, userService)).Methods(http.MethodPost)

	router.HandleFunc("/doc/new/{userId}", document.NewDocumentHandle(logger, docService, userService)).Methods(http.MethodPost)
	router.HandleFunc("/doc/get/user/{userId}", document.GetAllUSerDocumentsByUserIdHandle(logger, docService, userService)).Methods(http.MethodGet)
	router.HandleFunc("/doc/update/{docId}", document.UpdateDocumentHandle(logger, docService)).Methods(http.MethodPatch)
	router.HandleFunc("/doc/delete/{docId}", document.DeleteDocumentHandle(logger, docService)).Methods(http.MethodDelete)

	router.HandleFunc("/ticket/new/{userId}", ticket.NewTicketHandle(logger, ticketService, userService)).Methods(http.MethodPost)
	router.HandleFunc("/ticket/all", ticket.GetAllTicketsHandle(logger, ticketService)).Methods(http.MethodGet)
	router.HandleFunc("/tickets/get/all/{ticketId}", ticket.GetAllInfoFromTicketIdHandle(logger, ticketService)).Methods(http.MethodGet)
	router.HandleFunc("/ticket/passengers/{ticketId}", ticket.GetAllPassengersByTicketIdHandle(logger, ticketService)).Methods(http.MethodGet)
	router.HandleFunc("/ticket/update/{ticketId}", ticket.UpdateTicketHandle(logger, ticketService)).Methods(http.MethodPatch)
	router.HandleFunc("/ticket/delete/{ticketId}", ticket.DeleteTicketHandle(logger, ticketService)).Methods(http.MethodDelete)

	// swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
