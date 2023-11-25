package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(14.90))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.50))
	fmt.Println(math.Round(1.49))
	fmt.Println(math.Max(10, 9))
	fmt.Println(math.Min(10, 9))
}
