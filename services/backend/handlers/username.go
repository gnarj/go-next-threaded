package handlers

import (
	"encoding/json"
	"net/http"
)

type UsernameResponse struct {
	Username string `json:"username"`
}

func UsernameHandler(w http.ResponseWriter, r *http.Request) {
	resp := UsernameResponse{Username: "gnarj"}
	json.NewEncoder(w).Encode(resp)
}
