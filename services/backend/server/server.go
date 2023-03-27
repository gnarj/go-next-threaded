package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func StartServer(addr string, router http.Handler) {
	log.Printf("Server is available at http://%s", addr)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router)))
}
