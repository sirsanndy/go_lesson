package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("error app")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
