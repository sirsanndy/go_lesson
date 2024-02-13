package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	db := flag.String("database", "postgresql", "Choose your dbms")

	flag.Parse()
	fmt.Println(*host, *db)
}
