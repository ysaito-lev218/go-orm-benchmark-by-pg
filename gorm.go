package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

func GormRead(id int) {

	connString := "host=192.168.200.10 user=postgres password=postgres dbname=test sslmode=disable"
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	user := User{Id: id}
	d := db.Find(&user)
	if d.Error != nil {
		panic(err)
	}
}


