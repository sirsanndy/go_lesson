package main

import "fmt"

func getFullname() (string, string) {
	return "Sandy", "Hasanudin"
}

func main() {
	firstName, lastName := getFullname()

	onlyFirstName, _ := getFullname()

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println("Only Firstname :", onlyFirstName)
}
