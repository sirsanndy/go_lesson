package main

import "fmt"

func main() {
	var more_than = 16 > 14
	var less_than = 20 < 30
	var equal_more_than = 20 >= 20
	var equal = 30 == 30
	var not_equal = 30 != 10

	fmt.Println(more_than)
	fmt.Println(less_than)
	fmt.Println(equal_more_than)
	fmt.Println(equal)
	fmt.Println(not_equal)
}
