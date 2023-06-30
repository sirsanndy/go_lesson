package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var sandy Customer
	sandy.Name = "Sandy"
	sandy.Address = "Bogor"
	sandy.Age = 29

	fmt.Println(sandy.Name)
	fmt.Println(sandy.Address)
	fmt.Println(sandy.Age)
	fmt.Println(sandy)
}
