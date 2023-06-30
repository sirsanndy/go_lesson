package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	fmt.Println(getHello("Sandy"))
	fmt.Println(getHello(""))
}
