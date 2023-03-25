package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func TodosHandler(w http.ResponseWriter, db *sql.DB, r *http.Request) {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos.todos")
	if err != nil {
		log.Fatalln(err)
		// c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	defer rows.Close()
	json.NewEncoder(w).Encode(todos)
}
