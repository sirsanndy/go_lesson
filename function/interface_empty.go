package main

import "fmt"

func Ups() interface{} {
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
