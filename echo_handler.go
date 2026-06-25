package main

import (
	"butler-demo/shared"
	"encoding/json"
	"net/http"
)

type EchoResponse struct {
	Message string `json:"message"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// For simplicity, we'll read from query params or body
	// Using query param "input" for the echo value
	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Missing input parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := EchoResponse{Message: shared.Echo(input)}
	json.NewEncoder(w).Encode(resp)
}
