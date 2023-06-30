package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Unknown type")
	}
}
