package adapter

import (
	"sync"

	"1/internal/domain"
)

type EthereumAdapter struct {
	Endpoint     string
	currentBlock int
	addresses    map[string]bool
	transactions map[string][]domain.Transaction
	mu           sync.Mutex
}

func NewEthereumAdapter(endpoint string) *EthereumAdapter {
	return &EthereumAdapter{
		Endpoint:     endpoint,
		addresses:    make(map[string]bool),
		transactions: make(map[string][]domain.Transaction),
	}
}
