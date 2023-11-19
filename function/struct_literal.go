package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
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
}
