package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}
	total := sumAll(numbers...)

	fmt.Println(total)
}d
