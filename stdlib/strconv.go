package main

import (
	"fmt"
	"strconv"
)

func main() {
	bool, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(bool)
	} else {
		fmt.Println("Error", err.Error())
	}

	result, err := strconv.Atoi("12")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	binary := strconv.FormatInt(123, 2)
	fmt.Println(binary)

	strInt := strconv.Itoa(123)
	fmt.Println(strInt)
}
