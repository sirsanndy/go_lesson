package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhamad"
	names[1] = "Sandy"
	names[2] = "Hasanudin"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{90, 95, 98}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	fmt.Println(len(names))
}
