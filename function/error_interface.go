package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Tidak dapat membagi dengan angka nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var result, err = Pembagian(100, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error occurred :", err.Error())
	}
}
