package main

import (
	"context"
	"net/http"
	"smartway-test/api"
	"smartway-test/config"
	"smartway-test/database"
	docRepo "smartway-test/database/doc_repo"
	ticketRepo "smartway-test/database/ticket_repo"
	userRepo "smartway-test/database/user_repo"
	"smartway-test/service"

	"smartway-test/lib/pctx"

	"go.uber.org/zap"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}

	//err = database.DownMigration(pgDb)
	//if err != nil {
	//	panic(err)
	//}

	err = database.UpMigrations(pgDb)
	if err != nil {
		panic(err)
	}

	var (
		userRepository = userRepo.NewUserRepository(logger, pgDb)
		userService    = service.NewUserService(logger, userRepository)

		docRepository = docRepo.NewDocRepository(logger, pgDb)
		docService    = service.NewDocService(logger, docRepository)

		ticketRepository = ticketRepo.NewDocRepository(logger, pgDb)
		ticketService    = service.NewTicketService(logger, ticketRepository)

		server = api.NewServer(ctxProvider, logger, settings, userService, docService, ticketService)
	)

	return App{
		logger:   logger,
		settings: settings,
		server:   server,
	}
}

func (a App) Run() {
	go func() {
		_ = a.server.ListenAndServe()
	}()
	a.logger.Debugf("HTTP server started on %d", a.settings.Port)
}

func (a App) Stop(ctx context.Context) {
	_ = a.server.Shutdown(ctx)
	a.logger.Debugf("HTTP server stopped")
}
