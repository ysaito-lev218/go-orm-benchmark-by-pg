package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

func GorpRead(id int) {

	connString := "host=192.168.200.10 user=postgres password=postgres dbname=test sslmode=disable"

	db, err := sql.Open("postgres", connString)
	checkErr(err)
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	defer dbmap.Db.Close()

	tb := User{}
	if _, err := dbmap.Get(&tb, 1); err != nil {
		panic(err)
	}
}
