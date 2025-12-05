package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	// Set header
	w.Header().Set("Content-Type", "application/json")
	// Set Status
	w.WriteHeader(status)
	// Return json
	json.NewEncoder(w).Encode(data)
}
