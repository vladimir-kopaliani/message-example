package httpserver

import (
	"context"
	"net"
	"net/http"
	"time"
)

// HTTPServer represent http server
type HTTPServer struct {
	logger Logger
	server *http.Server
}

// Handler used in 'Configuration' to set handler for http server
type Handler struct {
	Path    string
	Handler http.Handler
}

// Configuration contains required settings for http server
type Configuration struct {
	Logger   Logger
	Address  string
	Handlers []Handler
}

// NewHTTPServer creates a new instance of http server
func NewHTTPServer(ctx context.Context, conf *Configuration) (*HTTPServer, error) {
	if conf == nil {
		panic("configuration is not set for http server")
	}
	if conf.Logger == nil {
		panic("logger is not set for http server")
	}
	if conf.Address == "" {
		conf.Logger.Fatal("address is not set for http server")
	}

	baseContext := func(net.Listener) context.Context {
		return ctx
	}

	s := &http.Server{
		Addr:         conf.Address,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
		BaseContext:  baseContext,
	}

	for i := range conf.Handlers {
		http.Handle(conf.Handlers[i].Path, conf.Handlers[i].Handler)
	}

	return &HTTPServer{
		logger: conf.Logger,
		server: s,
	}, nil
}

// Launch ...
func (s HTTPServer) Launch(ctx context.Context) error {
	s.logger.Info("HTTP is listening on " + s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Fatal(err)
		return err
	}

	return nil
}

// Close ...
func (s HTTPServer) Close(ctx context.Context) error {
	s.logger.Info("HTTP server is shutting down...")

	err := s.server.Shutdown(ctx)
	if err != nil {
		s.logger.Error(err)
		return nil
	}

	s.logger.Info("HTTP server is off.")

	return nil
}
