package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	sandy := Customer{"Sandy", "Bogor", 29}

	fmt.Println(sandy)

	sandy = Customer{
		Name:    "Sandy",
		Address: "Batu",
		Age:     29,
	}

	fmt.Println(sandy)
	sandy.sayHi("there")
}
