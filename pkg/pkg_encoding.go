package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var val string = "Muhamad Sandy Hasanudin"

	var encoded string = base64.StdEncoding.EncodeToString([]byte(val))
	fmt.Println(encoded)
	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
