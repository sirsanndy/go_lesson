package main

import "fmt"

func main() {
	var cetak = false

	var name = "Sandy"

	if cetak {
		fmt.Println(name)
	} else if cetak || name == "Sandy" {
		fmt.Println("Else If :", name)
	} else {
		fmt.Println("Nama tidak tercetak")
	}
}
