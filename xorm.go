package main

import "github.com/go-xorm/xorm"

func XormRead(id int) {

	connString := "host=192.168.200.10 user=postgres password=postgres dbname=test sslmode=disable"
	engine, _ := xorm.NewEngine("postgres", connString)

	db := engine.NewSession()
	if err := engine.Sync2(new(User)); err != nil {
		panic(err)
	}
	defer db.Close()

	user := User{}
	if _, err := db.Id(id).Get(&user); err != nil {
		panic(err)
	}
}
