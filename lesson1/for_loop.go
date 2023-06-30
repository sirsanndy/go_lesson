package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Count", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Idx", i)
	}

	slice := []string{"Muhamad", "Sandy", "Hasanudin"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Value [", i, "]", value)
	}

	person := map[string]string{
		"name":    "Muhamad Sandy Hasanudin",
		"address": "Bogor",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
