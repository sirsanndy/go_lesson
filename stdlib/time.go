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

	formatter := "Mon, 2006-01-02 15:04:05"

	var val = utc.Format(formatter)
	fmt.Println(val)
}
