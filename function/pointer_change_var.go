package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	sandy := Person{
		Name:    "Sandy",
		Age:     29,
		Address: "Bogor",
	}

	var sandyPointer *Person = &sandy
	sandyPointer.Name = "Sandy Pointer"

	sandyPointer = &Person{"Sandy 2", 29, "Bandung"}

	fmt.Println(sandy)
	fmt.Println(sandyPointer)
}
