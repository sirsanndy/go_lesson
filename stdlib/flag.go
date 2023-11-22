package main

import (
	"flag"
	"fmt"
)

func main() {
	var firstName = flag.String("firstname", "firstname", "write your first name")
	var lastName = flag.String("lastname", "lastname", "write your last name")

	flag.Parse()
	fmt.Println("Hello", *firstName, *lastName)
}
