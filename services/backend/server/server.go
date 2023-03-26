package server

import (
	"log"
	"net/http"
)

func StartServer(addr string, router http.Handler) (*http.Server, error) {
	log.Printf("Server is available at http://%s", addr)
	srv := &http.Server{Addr: addr, Handler: router}
	err := srv.ListenAndServe()
	if err != nil {
		return nil, err
	}
	return srv, nil
}
