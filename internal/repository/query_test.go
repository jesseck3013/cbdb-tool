package repository

import (
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
