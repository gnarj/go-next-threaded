package db

import (
	"testing"
)

func TestCreateDBConnection(t *testing.T) {
	db, err := CreateDBConnection()
	if err != nil {
		t.Fatalf("CreateDBConnection() failed with error: %v", err)
	}
	if db == nil {
		t.Errorf("expected a non-nil database connection")
	}
}
