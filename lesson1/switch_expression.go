package main

import "fmt"

func main() {
	name := "Iffy"
	switch name {
	case "Sandy":
		fmt.Println("Hello,", name)

	default:
		fmt.Println("Hello Stranger!")
	}

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("Lebih dari 3")

	case false:
		fmt.Println("Kurang dari sama dengan 3")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Lebih dari 10")

	case length < 10:
		fmt.Println("Kurang dari sama dengan 10")
	}
}
