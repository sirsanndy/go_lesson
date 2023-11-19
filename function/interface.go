package main

import "fmt"

type Person struct {
	Name string
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}
func main() {
	person := Person{Name: "Sandy"}
	SayHello(person)
}
