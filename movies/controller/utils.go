package controller

import (
	"encoding/json"
	"net/http"
)

// SendResponse - to send response back to client
func SendResponse(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// SendError - to send error response to client
func SendError(w http.ResponseWriter, err error) {
	response, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(response)
}
