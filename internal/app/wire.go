//go:build wireinject
// +build wireinject

package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"1/internal/config"
	"1/internal/service"
	httpServer "1/internal/transport/http"

	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		NewLogger,
		config.New,
		NewHTTPServerConfig,
		NewHTTPServerMux,
		NewHTTPServer,
		service.New,
		httpServer.NewServer,
		wire.Struct(new(App), "svc", "cfg", "logger", "server"),
	)
	return &App{}, nil
}

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "INFO: ", log.LstdFlags)
}

func NewHTTPServerConfig(cfg config.Config) httpServer.ServerConfig {
	return httpServer.ServerConfig{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
	}
}

func NewHTTPServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	httpServer.RegisterRoutes(mux)
	return mux
}

func NewHTTPServer(cfg httpServer.ServerConfig, mux *http.ServeMux, logger *log.Logger) *http.Server {
	return &http.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
}
