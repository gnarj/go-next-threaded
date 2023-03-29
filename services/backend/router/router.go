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
	r.HandleFunc("/add-todo", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTodoHandler(w, db, r)
	}).Methods("POST")
	r.HandleFunc("/update-todo", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTodoHandler(w, db, r)
	}).Methods("PUT")
	r.HandleFunc("/username", handlers.UsernameHandler).Methods("GET")

	return r
}
