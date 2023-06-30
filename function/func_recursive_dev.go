package main

import "fmt"

func recursive(value int) int {
	if value <= 0 {
		return 1
	}

	return value * recursive(value-1)
}

func main() {
	fmt.Println(recursive(5))
}
