package main

import (
	"backend/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register endpoints using handlers from separate files
	r.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
	r.HandleFunc("/username", handlers.UsernameHandler).Methods("GET")

	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
