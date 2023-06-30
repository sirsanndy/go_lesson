package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		counter++
	}
	increment()

	fmt.Println(counter)
}
