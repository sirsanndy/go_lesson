package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int16
}

func main() {
	var person *Person = new(Person)
	var person1 *Person = person

	person1.Name = "Sandy"
	fmt.Println(person)
	fmt.Println(person1)
}
