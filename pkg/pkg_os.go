package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)
	}
}
