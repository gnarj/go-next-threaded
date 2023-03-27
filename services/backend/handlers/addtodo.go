package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func AddTodoHandler(w http.ResponseWriter, db *sql.DB, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Extract the "todo" field from the JSON body
	var payload struct {
		Todo string `json:"todoText"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the todo into the database
	query := "INSERT INTO todos.todos (item) VALUES ($1)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing database statement: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, payload.Todo)
	if err != nil {
		log.Printf("Error executing database statement: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error getting number of rows affected: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%d products created ", rows)

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo added successfully"))
}
