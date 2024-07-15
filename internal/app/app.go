package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"1/internal/adapter/ethereum"
	"1/internal/config"
	"1/internal/service"
	httpServer "1/internal/transport/http"
)

func Run() error {
	cfg, err := config.New[config.Config]()
	if err != nil {
		return err
	}

	trustAdapter := ethereum.NewEthereumAdapter(cfg.Ethereum.Endpoint)
	trustService := service.New(cfg, trustAdapter)
	handler := httpServer.New(trustService)
	router := mux.NewRouter()
	httpServer.RegisterRoutes(router, handler)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port),
		Handler: router,
	}
	stopped := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port))

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Have a nice day!")

	return nil
}
