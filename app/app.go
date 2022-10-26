package app

import (
	"github.com/5verst/internal/store"
	"github.com/rs/zerolog"
)

type Application struct {
	Log   *zerolog.Logger
	Store store.Interface
}
