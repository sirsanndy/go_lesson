package main

import "fmt"

// ... variadic must be last arg
func sumAll(n int, numbers ...int) int {
	if n <= 0 {
		return 0
	} else {
		return sumAll(n-1, numbers...) + numbers[n-1]
	}
}

func main() {
	numbers := []int{10, 20, 30}
	total := sumAll(len(numbers), numbers...)

	fmt.Println(total)
}
