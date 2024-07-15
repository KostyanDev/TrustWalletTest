package http

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/subscribe", handler.Subscribe).Methods("GET")
	router.HandleFunc("/transactions", handler.GetTransactions).Methods("GET")
}
