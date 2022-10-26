package service

import (
	"github.com/5verst/app"
	"github.com/5verst/internal/store"
	"github.com/rs/zerolog"
)

type EventsService struct {
	log   *zerolog.Logger
	store store.Interface
}

func NewEventsService(a *app.Application) *EventsService {
	return &EventsService{
		log:   a.Log,
		store: a.Store,
	}
}
