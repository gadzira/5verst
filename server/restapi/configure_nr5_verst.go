// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/5verst/server/models"
	"github.com/5verst/server/restapi/operations"
	"github.com/5verst/server/restapi/operations/events"
	"github.com/5verst/server/restapi/operations/index"
)

//go:generate swagger generate server --target ../../server --name Nr5Verst --spec ../../spec/flatten.swagger.yaml --principal models.Auth --exclude-main

func configureFlags(api *operations.Nr5VerstAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Nr5VerstAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.JWTAuth == nil {
		api.JWTAuth = func(token string) (*models.Auth, error) {
			return nil, errors.NotImplemented("api key auth (JWT) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.IndexIndexHandler == nil {
		api.IndexIndexHandler = index.IndexHandlerFunc(func(params index.IndexParams) middleware.Responder {
			return middleware.NotImplemented("operation index.Index has not yet been implemented")
		})
	}
	if api.EventsCreateEventHandler == nil {
		api.EventsCreateEventHandler = events.CreateEventHandlerFunc(func(params events.CreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events.CreateEvent has not yet been implemented")
		})
	}
	if api.EventsListEventsHandler == nil {
		api.EventsListEventsHandler = events.ListEventsHandlerFunc(func(params events.ListEventsParams) middleware.Responder {
			return middleware.NotImplemented("operation events.ListEvents has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
