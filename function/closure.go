package main

import "fmt"

func main() {
	var counter = 0

	var increment = func() {
		counter++
		fmt.Println("Increased")
	}

	increment()
	increment()
	fmt.Println(counter)
}
