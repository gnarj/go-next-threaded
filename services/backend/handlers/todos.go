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

type TodoItem struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

func TodosHandler(w http.ResponseWriter, db *sql.DB, r *http.Request) {
	var todos []TodoItem
	rows, err := db.Query("SELECT id, item FROM todos.todos ORDER BY id asc")
	if err != nil {
		log.Fatalln(err)
		// c.JSON("An error occured")
	}
	for rows.Next() {
		var id int
		var item string
		err := rows.Scan(&id, &item)
		if err != nil {
			log.Fatalln(err)
		}
		todo := TodoItem{id, item}
		todos = append(todos, todo)
	}
	defer rows.Close()

	if len(todos) == 0 {
		todos = []TodoItem{}
	}

	json.NewEncoder(w).Encode(todos)
}

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
	log.Printf("%d todo(s) created ", rows)

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo added successfully"))
}

func DeleteTodoHandler(w http.ResponseWriter, db *sql.DB, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Extract the id field from the JSON body
	var payload struct {
		Id int `json:"id"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete the todo from the database
	query := "DELETE FROM todos.todos WHERE id = $1"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing database statement: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, payload.Id)
	log.Printf("Deleting todo with ID %d\n", payload.Id)

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
	log.Printf("%d todo deleted ", rows)

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo deleted successfully"))
}

func UpdateTodoHandler(w http.ResponseWriter, db *sql.DB, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Extract the item and id field from the JSON body
	var payload struct {
		Id   int    `json:"id"`
		Todo string `json:"item"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update the todo into the database
	query := "UPDATE todos.todos SET item = $1 WHERE id = $2"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing database statement: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, payload.Todo, payload.Id)
	log.Printf("Updating todo with ID %d and item '%s'\n", payload.Id, payload.Todo)

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
	log.Printf("%d todo updated ", rows)

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo updated successfully"))
}
