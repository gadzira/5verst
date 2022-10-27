package app

import (
	"context"

	"github.com/5verst/config"
	"github.com/5verst/internal/store"
	"github.com/rs/zerolog"
)

type Application struct {
	Context context.Context
	Log     *zerolog.Logger
	Config  *config.Config
	Store   store.Interface
}

func NewApplication(ctx context.Context, log *zerolog.Logger, conf *config.Config, st store.Interface) *Application {
	return &Application{
		Context: ctx,
		Log:     log,
		Config:  conf,
		Store:   st,
	}
}
