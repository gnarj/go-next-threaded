package router

import (
	"backend/handlers"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Register endpoints
	r.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
	r.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		handlers.TodosHandler(w, db, r)
	}).Methods("GET")
	r.HandleFunc("/username", handlers.UsernameHandler).Methods("GET")

	return r
}
