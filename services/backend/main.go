package main

import (
	"backend/db"
	"backend/handlers"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

func main() {
	db, err := db.CreateDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Register endpoints
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		handlers.StatusHandler(w, r)
	}).Methods("GET")

	r.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		handlers.TodosHandler(w, db, r)
	}).Methods("GET")

	r.HandleFunc("/username", handlers.UsernameHandler).Methods("GET")

	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
