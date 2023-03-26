package main

import (
	"backend/db"
	"backend/router"
	"backend/server"

	"log"

	_ "github.com/lib/pq"
)

const TestAddr = "localhost:8000"

func main() {
	// create database connection
	db, err := db.CreateDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	// create router
	r := router.NewRouter(db)

	// start server
	server.StartServer(TestAddr, r)
}
