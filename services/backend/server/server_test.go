package server

import (
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestStartServer(t *testing.T) {
	addr := "localhost:8000"
	router := mux.NewRouter()

	// Register test endpoints
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	}).Methods("GET")

	// Launch server in a goroutine
	go func() {
		StartServer(addr, router)
	}()
	log.Printf("Server started at %s", addr)

	// Make test request to server
	resp, err := http.Get("http://" + addr + "/test")
	if err != nil {
		t.Fatalf("Error making test request: %v", err)
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	log.Printf("Test request succeeded")
}
