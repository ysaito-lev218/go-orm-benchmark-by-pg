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

	tb := User{}

	connString := "host=192.168.200.10 user=postgres password=postgres dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//if _, err := db.Query(selectSQL, id); err != nil {
	//	checkErr(err)
	//}

	rows, err := db.Query(selectSQL, id)
	if err != nil {
		panic(err)
	}
	if !rows.Next() {
		panic(sql.ErrNoRows)
	}
	err = rows.Scan(&tb.Id,
		&tb.Address,
		&tb.Age,
		&tb.Name,
		&tb.Tel)

	if err != nil {
		panic(err)
	}
	rows.Close()
}
