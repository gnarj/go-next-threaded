package router

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

func TestNewRouter(t *testing.T) {
	// Create a new router with a mock database connection
	db := getMockDB()
	r := NewRouter(db)

	// Define expected routes
	expectedRoutes := []struct {
		Name   string
		Method string
		Path   string
	}{
		{"StatusHandler", "GET", "/status"},
		{"TodosHandler", "GET", "/todos"},
		{"UsernameHandler", "GET", "/username"},
	}

	// Verify that each expected route is present in the router
	for _, expected := range expectedRoutes {
		found := false
		err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			// Get information about the route
			method, _ := route.GetMethods()
			path, _ := route.GetPathTemplate()

			// Check if the route matches the expected route
			if method[0] == expected.Method && path == expected.Path {
				found = true
			}

			return nil
		})
		if err != nil {
			t.Errorf("Error while walking routes: %s", err)
		}
		if !found {
			t.Errorf("Expected route %s %s not found", expected.Method, expected.Path)
		}
	}
}

func getMockDB() *sql.DB {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockDB.Close()
	return mockDB
}
