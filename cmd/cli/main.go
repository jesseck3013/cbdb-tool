/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"database/sql"
	"log"

	"github.com/jesseck3013/cbdb-tool/cmd/cli/cmd"
	"github.com/jesseck3013/cbdb-tool/internal/repository"
	_ "modernc.org/sqlite"
)

func main() {
	// TODO: embed sqlite
	dsn := "file:./data/cbdb.sqlite3?mode=ro"
	sqlite, err := sql.Open("sqlite", dsn)

	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer sqlite.Close()

	db := repository.New(sqlite)
	cmd.Execute(db)
}
