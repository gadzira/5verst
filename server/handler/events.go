package handler

import (
	"github.com/5verst/app"
	"github.com/5verst/internal/service"
	"github.com/rs/zerolog"
)

type EventsHandler struct {
	log           *zerolog.Logger
	eventsService *service.EventsService
}

func NewEventsHandler(a *app.Application) *EventsHandler {
	return &EventsHandler{
		log:           a.Log,
		eventsService: service.NewEventsService(a),
	}
}
