package store

import (
	"context"

	"github.com/5verst/internal/domain"
)

type Interface interface {
	EventsInterface
}

type EventsInterface interface {
	LookupEvents(ctx context.Context, e domain.EventsLookupForm) (domain.EventsList, error)
	CreateEvents(ctx context.Context, e *domain.Event) (domain.EventsList, error)
}
