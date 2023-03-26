package server

import (
	"log"
	"net/http"
)

func StartServer(addr string, router http.Handler) {
	log.Printf("Server is available at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
