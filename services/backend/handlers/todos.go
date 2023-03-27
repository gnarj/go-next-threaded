package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

func TodosHandler(w http.ResponseWriter, db *sql.DB, r *http.Request) {
	var todos []TodoItem
	rows, err := db.Query("SELECT id, item FROM todos.todos")
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
	json.NewEncoder(w).Encode(todos)
}
