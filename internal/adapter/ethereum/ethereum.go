package ethereum

import (
	"context"

	"1/internal/domain"
)

func (e *EthereumAdapter) GetCurrentBlock(ctx context.Context) (int, error) {
	err := e.fetchCurrentBlock(ctx)
	if err != nil {
		return 0, err
	}
	return e.currentBlock, nil
}

func (e *EthereumAdapter) Subscribe(ctx context.Context, address string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, exists := e.addresses[address]; exists {
		return false
	}
	e.addresses[address] = true
	return true
}

func (e *EthereumAdapter) GetTransactions(ctx context.Context, address string) ([]domain.Transaction, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	transactions, err := e.fetchTransactions(ctx, address)
	if err != nil {
		return nil, err
	}

	e.transactions[address] = transactions

	return transactions, nil
}
