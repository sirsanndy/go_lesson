package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	fmt.Println(now)

	var utc = time.Date(2024, time.July, 8, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())
}
