package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jesseck3013/cbdb-tool/internal/repository"
	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite", "../../data/cbdb.sqlite3")
	if err != nil {
		panic(err)
	}

	queries := repository.New(db)
	person, err := queries.GetPersonByID(ctx, 55870)
	if err != nil {
		panic(err)
	}

	fmt.Println(person.CPersonid, *person.CNameChn)
}
