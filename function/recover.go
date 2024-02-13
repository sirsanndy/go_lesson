package main

import "fmt"

func endApl() {
	message := recover()

	if message != nil {
		fmt.Println("Error dengan message :", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApl(error bool) {
	defer endApl()

	if error {
		panic("error app")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApl(true)
	fmt.Println("Hello")
}
