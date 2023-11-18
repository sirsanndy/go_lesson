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

	sandyPointer := &sandy
	sandy2 := &sandy
	sandyPointer.Name = "Sandy Pointer"

	fmt.Println(sandy)
	fmt.Println(sandyPointer)
	fmt.Println(sandy2)
}
