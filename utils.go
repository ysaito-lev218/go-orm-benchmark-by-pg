package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"fmt"
	"flag"
)

var (
	user          = flag.String("user", "postgres", "the database user")
	password      = flag.String("password", "postgres", "the database password")
	port *int     = flag.Int("port", 5432, "the database port")
	server        = flag.String("server", "192.168.200.10", "the database server")
	database      = flag.String("database", "test", "the database")

	connString string
)

// grom, pg
type User struct {
	Id      int
	Name    string
	Address string
	Tel     string
	Age     string
}

// genmai
type Users struct {
	Id      int
	Name    string
	Address string
	Tel     string
	Age     string
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func initDB() {

	sqls := []string{
		`DROP TABLE IF EXISTS users;`,
		`CREATE TABLE users (
			id SERIAL NOT NULL,
			name text NOT NULL,
			address text NOT NULL,
			tel text NOT NULL,
			age integer NOT NULL,
			CONSTRAINT users_pkey PRIMARY KEY (id)
		) WITH (OIDS=FALSE);`,
	}

	DB, err := sql.Open("postgres", connString)
	checkErr(err)
	defer DB.Close()

	err = DB.Ping()
	checkErr(err)

	for _, sql := range sqls {
		_, err = DB.Exec(sql)
		checkErr(err)
	}

	insert :=
	`INSERT INTO users (
		name,
		address,
		tel,
		age
	) VALUES (
		'ORM benchmark',
		'Tokyo, Japan',
		90012345678,
		100
	)`;

	for i := 1; i < 100000; i++ {
		_, err = DB.Exec(insert)
		checkErr(err)
	}
}

func initDB2() {

	sqls := []string{
		`DROP TABLE IF EXISTS test;`,
		`CREATE TABLE test (
			id SERIAL NOT NULL,
			name text NOT NULL,
			address text NOT NULL,
			tel text NOT NULL,
			age integer NOT NULL,
			CONSTRAINT test_pkey PRIMARY KEY (id)
		) WITH (OIDS=FALSE);`,
	}

	DB, err := sql.Open("postgres", connString)
	checkErr(err)
	defer DB.Close()

	err = DB.Ping()
	checkErr(err)

	for _, sql := range sqls {
		_, err = DB.Exec(sql)
		checkErr(err)
	}

	insert :=
	`INSERT INTO test (
		name,
		address,
		tel,
		age
	) VALUES (
		'ORM benchmark',
		'Tokyo, Japan',
		90012345678,
		100
	)`;

	for i := 1; i < 100000; i++ {
		_, err = DB.Exec(insert)
		checkErr(err)
	}
}
