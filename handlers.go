package main

import (
	"butler-demo/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := Response{Message: "Hello World!fhwiefhwoife"}
	json.NewEncoder(w).Encode(resp)
}

// handlers.go
// Add new endpoint “/answer” that should return {“message”: “The answer to the great question is 42”}
func answerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	answer := shared.GetAnswer()
	resp := Response{Message: fmt.Sprintf("The answer to the great question is %d", answer)}
	json.NewEncoder(w).Encode(resp)
}
