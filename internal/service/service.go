package service

import (
	"context"

	"1/internal/domain"

	"1/internal/config"
)

type EthereumAdapter interface {
	GetCurrentBlock(ctx context.Context) (int, error)
	Subscribe(ctx context.Context, address string) bool
	GetTransactions(ctx context.Context, address string) ([]domain.Transaction, error)
}

type Service struct {
	cfg             config.Config
	ethereumAdapter EthereumAdapter
}

func New(cfg config.Config, ethereumAdapter EthereumAdapter) *Service {
	return &Service{
		cfg:             cfg,
		ethereumAdapter: ethereumAdapter,
	}
}
