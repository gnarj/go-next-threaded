package handlers

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	resp := StatusResponse{Status: "ok"}
	json.NewEncoder(w).Encode(resp)
}
