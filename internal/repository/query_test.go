package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func TestGetPersonByID(t *testing.T) {
	// 1. The variable is named 'db'
	db, err := sql.Open("sqlite", "../../data/cbdb.sqlite3")
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	// 2. New() belongs to the current package 'repository', so you call it directly.
	// You pass the 'db' variable cleanly without any shadowing conflicts.
	queries := New(db)

	person, err := queries.GetPersonByID(context.Background(), 1762)

	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}

	if *person.CNameChn != "王安石" {
		t.Errorf("Expected '王安石', got '%s'", *person.CNameChn)
	}
}
