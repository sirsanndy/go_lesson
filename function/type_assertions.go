package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)

	fmt.Println(resultString)
}
