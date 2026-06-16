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
	// TODO: handle the DB path
	db, err := sql.Open("sqlite", "./data/cbdb.sqlite3")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	query := repository.New(db)

	cmd.Execute(query)
}
