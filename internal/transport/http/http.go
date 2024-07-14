package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	logger *log.Logger
}

type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(server *http.Server, logger *log.Logger) *Server {
	return &Server{
		server: server,
		logger: logger,
	}
}

func (s *Server) ListenAndServe() error {
	s.logger.Printf("Listening on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
