package main

import (
	"context"
	"fmt"

	"github.com/5verst/config"
	"github.com/5verst/internal/app"
	"github.com/5verst/internal/server/http"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	a := app.NewApplication(ctx, log, &cfg, &db)

	server := http.NewServer(log, a)
	err := server.Start(a.Context, fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		log.Panic().Msg("run server")
	}
}
