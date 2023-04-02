package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Set up database connection information
	dbConfig := "host=localhost user=postgres dbname=todos port=5432 sslmode=disable"

	// Check if a seed file exists
	seedPath := filepath.Join("migrations", "seed_data.sql")
	_, err := os.Stat(seedPath)
	if os.IsNotExist(err) {
		// Create a new seed file
		fmt.Println("Creating new seed file:", seedPath)
		if err := createSeedFile(); err != nil {
			fmt.Println("Error creating seed file:", err)
			return
		}
	} else {
		// Apply existing seed file
		fmt.Println("Applying existing seed file:", seedPath)
		if err := applySeedFile(dbConfig); err != nil {
			fmt.Println("Error applying seed file:", err)
			return
		}
	}

	// Apply database migrations
	fmt.Println("Applying database migrations")
	if err := applyMigrations(dbConfig); err != nil {
		fmt.Println("Error applying migrations:", err)
		return
	}

	fmt.Println("Database migration and seed data applied successfully!")
}

func createSeedFile() error {
	// Open a connection to the database
	connStr := "postgresql://postgres@localhost:5432/todos?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Query the database's information schema to get a list of all tables
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='todos'")
	if err != nil {
		return fmt.Errorf("error retrieving table names from database: %v", err)
	}
	defer rows.Close()

	// Create a new seed file with a timestamped name
	timestamp := time.Now().Format("20060102150405")
	path := fmt.Sprintf("./%s_seed.sql", timestamp)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating seed file: %v", err)
	}
	defer file.Close()

	// Loop through each table and generate seed data
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("error scanning table name: %v", err)
		}

		// Query the table and retrieve the data
		tableRows, err := db.Query(fmt.Sprintf("SELECT * FROM todos.%s", tableName))
		if err != nil {
			return fmt.Errorf("error retrieving data from table %s: %v", tableName, err)
		}
		defer tableRows.Close()

		// Write the data to the seed file in SQL format
		for tableRows.Next() {
			// Get a list of column names
			columns, err := tableRows.Columns()
			if err != nil {
				return fmt.Errorf("error retrieving column names for table %s: %v", tableName, err)
			}

			// Get a list of pointers to the values for each column
			values := make([]interface{}, len(columns))
			for i := range values {
				values[i] = new(interface{})
			}

			// Scan the row into the value pointers
			if err := tableRows.Scan(values...); err != nil {
				return fmt.Errorf("error scanning row in table %s: %v", tableName, err)
			}

			// Write the row to the seed file in SQL format
			valuesStrings := make([]string, len(columns))
			for i, v := range values {
				switch x := v.(type) {
				case nil:
					valuesStrings[i] = "NULL"
				case []byte:
					valuesStrings[i] = string(x)
				default:
					// Assert that v is a pointer and format it as a SQL string literal
					ptr, ok := v.(*interface{})
					if !ok {
						return fmt.Errorf("error converting value to pointer in table %s column %s", tableName, columns[i])
					}
					valuesStrings[i] = fmt.Sprintf("'%v'", *ptr)
				}
			}
			if _, err := file.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);\n", tableName, strings.Join(columns, ", "), strings.Join(valuesStrings, ", "))); err != nil {
				return fmt.Errorf("error writing to seed file: %v", err)
			}
		}

		// Check for any errors encountered while iterating over the rows
		if err := tableRows.Err(); err != nil {
			return fmt.Errorf("error iterating over table %s rows: %v", tableName, err)
		}
	}

	fmt.Println("Seed file created successfully:", path)
	return nil
}

func applySeedFile(dbConfig string) error {
	// Apply the seed file to the database using the "goose" migration tool
	command := exec.Command("./goose", "postgres", dbConfig, "seed")
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error applying seed data: %v\n%s", err, string(output))
	}
	fmt.Println(string(output))
	return nil
}

func applyMigrations(dbConfig string) error {
	// Apply all unapplied migrations to the database using the "goose" migration tool
	command := exec.Command("./goose", "postgres", dbConfig, "up")
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error applying migrations: %v\n%s", err, string(output))
	}
	fmt.Println(string(output))
	return nil
}
