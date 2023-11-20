package main

import (
	"fmt"
	"function/helper"
)

func main() {
	var result = helper.SayHello("Sandy")
	fmt.Println(helper.App)
	fmt.Println(result)
}
