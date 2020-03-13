package server

import (
	"encoding/json"
	"net/http"
)

// responseBody defines the structure of the response back
// to caller.
type responseBody struct {
	IP string `json:"ip"`
}

// jsonResponse performs JSON encoding and writes to the provided
// http.ResponseWriter.
func jsonResponse(w http.ResponseWriter, s string) {
	if err := json.NewEncoder(w).Encode(&responseBody{IP: s}); err != nil {
		panic(err)
	}
}

// textResponse writes the plain text as bytes to the provided
// http.ResponseWriter.
func textResponse(w http.ResponseWriter, s string) {
	w.Write([]byte(s))
}
