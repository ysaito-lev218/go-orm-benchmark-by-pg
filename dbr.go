package main

import (
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
)

func DbrRead(id int) {

	connString := "host=192.168.200.10 user=postgres password=postgres dbname=test sslmode=disable"

	conn, _ := dbr.Open("postgres", connString, nil)
	db := conn.NewSession(nil)
	defer db.Close()

	//var user []User
	//if _, err := db.Select("*").From("user").Load(&user); err != nil {
	//if _, err := db.Select("*").From("user").Where("Id = ?", id).Load(&user); err != nil {
	//var user User
	//if _, err := db.Select("*").From("user").Where("id = ?", id).Load(&user); err != nil {
	//	panic(err)
	//}

	var user User
	db.Select("*").From("user").Where("id = ?", 1).LoadStruct(&user)
}
