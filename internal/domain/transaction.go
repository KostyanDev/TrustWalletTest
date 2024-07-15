package domain

type Transaction struct {
	BlockNumber string `json:"blockNumber"`
	Hash        string `json:"hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
}
