package http

import (
	"context"

	"1/internal/service"

	"1/internal/domain"
)

type Service interface {
	GetCurrentBlockService(ctx context.Context) (int, error)
	SubscribeService(ctx context.Context, address string) bool
	GetTransactionsService(ctx context.Context, address string) ([]domain.Transaction, error)
}

type Handler struct {
	service Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
