package http

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}
	if h.service.SubscribeService(r.Context(), address) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Subscribed to " + address))
	} else {
		http.Error(w, "already subscribed", http.StatusBadRequest)
	}
}

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}
	transactions, err := h.service.GetTransactionsService(r.Context(), address)
	if err != nil {
		http.Error(w, "failed to get transactions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(transactions)
}
