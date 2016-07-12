package main

import (
	"flag"
	"runtime"
	"fmt"
)

func main()  {

	flag.Parse()

	runtime.GOMAXPROCS(1 * runtime.NumCPU())

	connString = fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", *server, *user, *password, *port, *database)

	initDB()
}
