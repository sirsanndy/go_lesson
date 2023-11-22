package main

import (
	"fmt"
	"os"
)

func SayHello(names []string) string {
	var fullName string
	for _, name := range names {
		fullName += name + " "
	}
	return "Hello " + fullName
}

func main() {
	var args = os.Args
	fmt.Println(SayHello(args))
}
