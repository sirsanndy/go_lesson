package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int16
}

func changeName(person *Person, name string) {
	person.Name = name
}

func main() {
	var person Person = Person{"Sandy", "Batu", 29}
	changeName(&person, "Ilham")
	fmt.Println(person)
}
