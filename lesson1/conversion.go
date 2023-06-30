package main

import "fmt"

func main() {
	var nilai int16 = 3278

	var nilai32 int32 = int32(nilai)
	var nilai8 int8 = int8(nilai)

	fmt.Println(nilai32)
	fmt.Println(nilai8)

	var name = "Sandy"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(eString)
}
