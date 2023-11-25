package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Sandy Hasanudin", "Sandy"))
	fmt.Println(strings.Split("Sandy Hasanudin", " "))
	fmt.Println(strings.ToLower("Sandy Hasanudin"))
	fmt.Println(strings.ToUpper("Sandy Hasanudin"))
	fmt.Println(strings.Trim("   Sandy Hasanudin   ", " "))
	fmt.Println(strings.ReplaceAll("M Sandy Hasanudin", "M", "Muhamad"))
}
