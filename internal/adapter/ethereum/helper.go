package ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"1/internal/domain"
)

func (e *EthereumAdapter) sendRequest(ctx context.Context, payload map[string]interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, e.Endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	return resp, nil
}

func (e *EthereumAdapter) fetchCurrentBlock(ctx context.Context) error {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []interface{}{},
		"id":      1,
	}

	resp, err := e.sendRequest(ctx, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	blockHex := result["result"].(string)
	var blockNumber int
	fmt.Sscanf(blockHex, "%x", &blockNumber)
	e.currentBlock = blockNumber
	return nil
}

func (e *EthereumAdapter) fetchTransactions(ctx context.Context, address string) ([]domain.Transaction, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getLogs",
		"params": []interface{}{
			map[string]interface{}{
				"address":   address,
				"fromBlock": "0x1034500", // 16999000 в шестнадцатеричном формате
				"toBlock":   "0x103462a",
			},
		},
		"id": 1,
	}

	log.Printf("Sending request to Ethereum node: %v", payload)
	resp, err := e.sendRequest(ctx, payload)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Result []struct {
			BlockNumber string `json:"blockNumber"`
			Hash        string `json:"transactionHash"`
			From        string `json:"from"`
			To          string `json:"to"`
			Value       string `json:"value"`
		} `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	var transactions []domain.Transaction
	for _, tx := range result.Result {
		transactions = append(transactions, domain.Transaction{
			BlockNumber: tx.BlockNumber,
			Hash:        tx.Hash,
			From:        tx.From,
			To:          tx.To,
			Value:       tx.Value,
		})
	}

	log.Printf("Fetched %d transactions for address %s", len(transactions), address)
	return transactions, nil
}
