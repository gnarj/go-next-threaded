package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUsernameHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/username", nil)
	if err != nil {
		t.Fatalf("could not create http request: %v", err)
	}

	rec := httptest.NewRecorder()

	UsernameHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", rec.Code, http.StatusOK)
	}

	want := UsernameResponse{Username: "gnarj"}
	var got UsernameResponse

	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("could not decode the response body: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected response body: got %v, want %v", got, want)
	}
}
