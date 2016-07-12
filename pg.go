package main

import (
	"gopkg.in/pg.v4"
)

func PgRead(id int) {

	db := pg.Connect(&pg.Options{
		Addr: "192.168.200.10:5432",
		User: "postgres",
		Password: "postgres",
		Database: "test",
	})
	defer db.Close()

	user := User{Id: id}
	if err := db.Select(&user); err != nil {
		panic(err)
	}
}
