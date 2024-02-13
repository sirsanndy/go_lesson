package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := ring.New(3)

	data.Value = "Muhamad"
	data = data.Next()
	data.Value = "Hasanudin"

	data = data.Prev()
	data.Value = "Sandy"

	data.Do(func(val any) {
		fmt.Println(val)
	})
}
