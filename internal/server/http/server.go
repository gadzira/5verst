package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/5verst/internal/app"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Application interface{}

type Server struct {
	router *mux.Router
	log    *zerolog.Logger
	app    *app.Application
}

func NewServer(log *zerolog.Logger, app *app.Application) *Server {
	return &Server{
		log: log,
		app: app,
	}
}

func (s *Server) Start(ctx context.Context, addr string) error {
	s.initializeRoutes()
	s.log.Info().Msgf("server run up on %s", addr)

	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		s.log.Panic().Msgf("run server: %v", err)
		return fmt.Errorf("run server: %v", err)
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}

func (s *Server) initializeRoutes() {
	router := mux.NewRouter()
	router.Handle("/list_of_events", LookupEvents(s.app)).Methods("GET")
	router.Handle("/create_new_event", CreateNewEvent(s.app)).Methods("POST")
}

func LookupEvents(a *app.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		fmt.Println(params)
	})
}

func CreateNewEvent(a *app.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		fmt.Println(body)
	})
}
