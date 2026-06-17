package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

const DB_PATH = "../../data/cbdb.sqlite3"

func ConnectDB(t *testing.T, path string) *Queries {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatalf("Failed to open DB: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return New(db)
}

func TestGetPersonByID(t *testing.T) {
	query := ConnectDB(t, DB_PATH)
	person, err := query.GetPersonByID(context.Background(), 1762)

	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}

	if *person.CNameChn != "王安石" {
		t.Errorf("Expected '王安石', got '%s'", *person.CNameChn)
	}
}
