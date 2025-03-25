//go:build mage

package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "words.db"

// InitDB initializes the SQLite database
func InitDB() error {
	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {
		fmt.Println("Database already exists")
		return nil
	}

	file, err := os.Create(dbPath)
	if err != nil {
		return fmt.Errorf("error creating database file: %v", err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	fmt.Println("Database initialized successfully")
	return nil
}

// Migrate runs database migrations
func Migrate() error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	migrationsDir := "db/migrations"
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("error reading migrations directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}

		content, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
		if err != nil {
			return fmt.Errorf("error reading migration file %s: %v", file.Name(), err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("error executing migration %s: %v", file.Name(), err)
		}

		fmt.Printf("Executed migration: %s\n", file.Name())
	}

	return nil
}

// Status shows database table information
func Status() error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Query to get table names
	rows, err := db.Query(`
		SELECT name FROM sqlite_master 
		WHERE type='table' 
		ORDER BY name;
	`)
	if err != nil {
		return fmt.Errorf("error querying tables: %v", err)
	}
	defer rows.Close()

	fmt.Println("Database Tables:")
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		fmt.Printf("- %s\n", name)
	}

	return nil
}
