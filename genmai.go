package main

import (
	_ "github.com/lib/pq"
	"github.com/naoina/genmai"
)

func GenmaiRead(id int) {

	connString := "host=192.168.200.10 user=postgres password=postgres dbname=test sslmode=disable"
	db, err := genmai.New(&genmai.PostgresDialect{}, connString)
	checkErr(err)
	defer db.Close()

	var results []Users
	//if err := db.Select(&results); err != nil {
	if err := db.Select(&results, db.Where("id", "=", 1)); err != nil {
		panic(err)
	}
}
