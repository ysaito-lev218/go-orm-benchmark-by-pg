package main

import (
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
)

func DbrRead(id int) {

	connString := "host=localhost user=postgres password=postgres dbname=test sslmode=disable"

	conn, _ := dbr.Open("postgres", connString, nil)
	db := conn.NewSession(nil)
	defer db.Close()

	var user []User
	if _, err := db.Select("*").From("user").Load(&user); err != nil {
		panic(err)
	}
}
