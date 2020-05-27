package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-multierror"

	// pprof is imported as unused, but is required for debug profiling
	_ "net/http/pprof"
)

type Server struct {
	config    Config
	stdoutLog *log.Logger
}

type Config struct {
	Listen       string
	Port         string
	MaxBodySize  int64
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	DebugPerf    bool
}

func (config Config) Validate() error {
	var result *multierror.Error

	if config.Listen == "" {
		result = multierror.Append(result, errors.New("server listen address must be provided"))
	}
	if config.Port == "" {
		result = multierror.Append(result, errors.New("server port must be provided"))
	}
	if config.MaxBodySize == 0 {
		result = multierror.Append(result, errors.New("max body size must be provided and be greater than zero"))
	}

	return result.ErrorOrNil()
}

func New(config Config, stdoutLog *log.Logger) (*Server, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &Server{
		config:    config,
		stdoutLog: stdoutLog,
	}, nil
}

func (server *Server) Start() error {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range server.Routes() {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	if server.config.DebugPerf {
		router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	}

	server.stdoutLog.Printf("Starting server on %s:%s", server.config.Listen, server.config.Port)

	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", server.config.Listen, server.config.Port),
		ReadTimeout:  server.config.ReadTimeout,
		WriteTimeout: server.config.WriteTimeout,
		IdleTimeout:  server.config.IdleTimeout,
		Handler:      handlers.LoggingHandler(server.stdoutLog.Writer(), router),
	}

	return httpServer.ListenAndServe()
}
