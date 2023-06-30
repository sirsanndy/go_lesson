package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Muhamad Sandy Hasanudin",
		"address": "Bogor",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["department"] = "Information Technology"

	fmt.Println(person["department"])
	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Effective Java"
	book["author"] = "Joshua M. Bloch"
	book["price"] = "$15"

	fmt.Println(book)

	delete(book, "price")

	fmt.Println(book)
}
