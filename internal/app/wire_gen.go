// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"fmt"
	"log"
	http2 "net/http"
	"os"

	"1/internal/config"
	"1/internal/service"
	"1/internal/transport/http"
)

// Injectors from wire.go:

func InitializeApp() (*App, error) {
	logger := NewLogger()
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	serviceService := service.New(logger, configConfig)
	serverConfig := NewHTTPServerConfig(configConfig)
	serveMux := NewHTTPServerMux()
	server := NewHTTPServer(serverConfig, serveMux, logger)
	httpServer := http.NewServer(server, logger)
	app := &App{
		svc:    serviceService,
		cfg:    configConfig,
		logger: logger,
		server: httpServer,
	}
	return app, nil
}

// wire.go:

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "INFO: ", log.LstdFlags)
}

func NewHTTPServerConfig(cfg config.Config) http.ServerConfig {
	return http.ServerConfig{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
	}
}

func NewHTTPServerMux() *http2.ServeMux {
	mux := http2.NewServeMux()
	http.RegisterRoutes(mux)
	return mux
}

func NewHTTPServer(cfg http.ServerConfig, mux *http2.ServeMux, logger *log.Logger) *http2.Server {
	return &http2.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
}
