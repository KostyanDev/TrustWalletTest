package app

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"1/internal/config"
	"1/internal/service"
	httpServer "1/internal/transport/http"
)

type App struct {
	svc    *service.Service
	cfg    config.Config
	logger *log.Logger
	server *httpServer.Server
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Println("Starting application")

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		a.logger.Printf("Starting HTTP server on host,port - %s:%d\n", a.cfg.HTTPServer.Host, a.cfg.HTTPServer.Port)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatalf("Could not listen on host,port - %s,%d: %v\n", a.cfg.HTTPServer.Host, a.cfg.HTTPServer.Port, err)
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		a.logger.Println("Shutting down HTTP server...")

		ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := a.server.Shutdown(ctxShutDown); err != nil {
			a.logger.Fatalf("HTTP server Shutdown Failed:%+v", err)
			return err
		}

		a.logger.Println("HTTP server stopped")
		return nil
	})

	defer func() {
		a.logger.Println("Application stopped")
	}()

	return group.Wait()
}
