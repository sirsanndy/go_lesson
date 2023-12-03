package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < 5; i++ {
		var val string
		fmt.Print("Input value : ")
		_, err := fmt.Scanln(&val)

		if err != nil {
			fmt.Println("Error :", err)
		} else {
			data.Value = val
			data = data.Next()
		}
	}

	fmt.Println("\nPrint Circular List : ")

	for i := 0; i < data.Len(); i++ {
		if i < data.Len()-1 {
			fmt.Print(data.Value, ",")
		} else {
			fmt.Print(data.Value, "\n")
		}
		data = data.Next()
	}
}
