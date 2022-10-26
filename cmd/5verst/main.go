package main

import (
	"github.com/5verst/config"
	"github.com/5verst/pkg/database"
	"github.com/5verst/pkg/logger"
	"github.com/5verst/pkg/store"
	"github.com/rs/zerolog"
)

var log *zerolog.Logger

func main() {
	cfg := config.MustConfig()
	log := logger.MustLogger(cfg)
	connDB := database.MustConnectDB(cfg)
	db := store.MustDBStore(connDB)

	a := app.NewApplication(log, &db, &cfg)
	server := internal.NewHTTPServer(log, a)
}
