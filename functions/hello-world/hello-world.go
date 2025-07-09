package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

// Exported handler function expected by Netlify
func Handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Hello, World from Go on Netlify!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
