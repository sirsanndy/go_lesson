package main

import "fmt"

func main() {
	var firstName = "Sandy"
	var lastName = "Hasanudin"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%[2]s %[1]s'\n", firstName, lastName)
}
