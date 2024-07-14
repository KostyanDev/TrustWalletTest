package http

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handleRoot)
}

// handleRoot обработчик для корневого маршрута
func handleRoot(w http.ResponseWriter, r *http.Request) {
	defer handleError(w)
}

// handleError обрабатывает ошибки и отправляет сообщение об ошибке в ответ
func handleError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", r), http.StatusInternalServerError)
	}
}
