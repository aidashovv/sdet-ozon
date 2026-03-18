package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type HTTPServer struct {
	server *http.Server
	config Config
}

func NewHTTPServer(
	cfg Config,
	setupHandler http.Handler, getCbrXMLHandler http.Handler, deleteHandler http.Handler,
) *HTTPServer {
	r := chi.NewRouter()

	r.Post("/setup", setupHandler.ServeHTTP)
	r.Get("/scripts/XML_daily.asp", getCbrXMLHandler.ServeHTTP)

	return &HTTPServer{
		config: cfg,
		server: &http.Server{
			Addr:         cfg.Addr,
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.config.ShutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
