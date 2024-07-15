package service

import (
	"context"

	"1/internal/domain"
)

func (s *Service) GetCurrentBlockService(ctx context.Context) (int, error) {
	return s.ethereumAdapter.GetCurrentBlock(ctx)
}

func (s *Service) SubscribeService(ctx context.Context, address string) bool {
	return s.ethereumAdapter.Subscribe(ctx, address)
}

func (s *Service) GetTransactionsService(ctx context.Context, address string) ([]domain.Transaction, error) {
	return s.ethereumAdapter.GetTransactions(ctx, address)
}
