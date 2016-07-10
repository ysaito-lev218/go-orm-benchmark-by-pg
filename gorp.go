package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

func GorpRead(id int) {
	connString := "host=localhost user=postgres password=postgres dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {

	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	defer dbmap.Db.Close()

	var users []User
	//err = dbmap.SelectOne(&p2, "select * from posts where post_id=?", p2.Id)
	if _, err = dbmap.Select(&users, "select * from users"); err != nil {
		panic(err)
	}
}
