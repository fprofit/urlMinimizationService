package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func dbw() string {

	connString := "user=postgres password=MiX7681726 host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
		defer db.Close()
		fmt.Println(db)
	}

	return "ok"
}
