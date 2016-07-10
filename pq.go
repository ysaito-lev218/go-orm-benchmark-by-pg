package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	selectSQL = `SELECT id, name, address, tel, age FROM users WHERE id = $1`
)

func PqRead(id int) {

	connString := "host=localhost user=postgres password=postgres dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Query(selectSQL, id); err != nil {
		checkErr(err)
	}
}
