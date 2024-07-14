package service

import (
	"log"

	"1/internal/config"
)

type Service struct {
	logger    *log.Logger
	cfg       config.Config
	isRunning bool
}

func New(logger *log.Logger, cfg config.Config) *Service {
	logger.Println("Initializing service...")
	return &Service{
		logger:    logger,
		cfg:       cfg,
		isRunning: true,
	}
}
