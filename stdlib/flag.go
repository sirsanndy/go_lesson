package main

import (
	"flag"
	"fmt"
)

func SayHello(firstName *string, lastName *string) {
	fmt.Println("Hello", *firstName, *lastName)
}

func main() {
	var firstName = flag.String("firstname", "firstname", "write your first name")
	var lastName = flag.String("lastname", "lastname", "write your last name")

	flag.Parse()
	SayHello(firstName, lastName)
}
