package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatalf("could not create HTTP request: %v", err)
	}

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Call the StatusHandler function with the HTTP request and response recorder
	StatusHandler(rec, req)

	// Check the response status code
	if rec.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", rec.Code, http.StatusOK)
	}

	// Check the response body
	want := StatusResponse{Status: "ok"}
	var got StatusResponse

	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected response body: got %v, want %v", got, want)
	}
}
