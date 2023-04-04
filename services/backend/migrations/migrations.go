package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	migrationsDir = "./"
)

func main() {
	// Establish database connection
	db, err := sql.Open("postgres", "host=localhost user=postgres dbname=todos port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check if migration file exists
	migrationName := "update_database"
	migrationFilename := fmt.Sprintf("%s/%s.sql", migrationsDir, migrationName)
	if _, err := os.Stat(migrationFilename); os.IsNotExist(err) {
		// If migration file does not exist, create it with current schema and exit
		fmt.Println("Creating migration file...")
		if err := createMigrationFile(db, migrationFilename); err != nil {
			panic(err)
		}
		fmt.Printf("Migration file created: %s\n", migrationFilename)
		return
	}

	// If migration file exists, check for differences between migration and current schema
	fmt.Println("Checking for differences between migration file and current schema...")
	if err := checkForDifferences(db, migrationFilename); err != nil {
		// If differences found, migrate to the latest version
		if err := goose.Up(db, migrationsDir); err != nil {
			panic(err)
		}
		fmt.Println("Migrations applied successfully")
	} else {
		fmt.Println("No differences found")
	}
}

func createMigrationFile(db *sql.DB, migrationFilename string) error {
	// Generate SQL dump of current database schema
	dumpCmd := exec.Command("pg_dump", "-s", "-h", "localhost", "-U", "postgres", "-d", "todos")
	dumpOutput, err := dumpCmd.Output()
	if err != nil {
		return err
	}

	// Write SQL dump to migration file
	if err := ioutil.WriteFile(migrationFilename, []byte("-- +goose Up\n"+string(dumpOutput)+"\n-- +goose Down\n\n"), 0644); err != nil {
		return err
	}

	return nil
}

func checkForDifferences(db *sql.DB, migrationFilename string) error {
	// Generate SQL dump of current database schema
	dumpCmd := exec.Command("pg_dump", "-s", "-h", "localhost", "-U", "postgres", "-d", "todos")
	dumpOutput, err := dumpCmd.Output()
	if err != nil {
		return err
	}

	// Compare migration file with current schema
	migrationFile, err := ioutil.ReadFile(migrationFilename)
	if err != nil {
		return err
	}

	if string(dumpOutput) != strings.SplitN(string(migrationFile), "\n-- +goose Down\n", 2)[0][len("-- +goose Up\n"):len(string(migrationFile))] {
		return fmt.Errorf("differences found between migration file and current schema")
	}

	return nil
}
