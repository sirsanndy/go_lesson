package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Now()
	fmt.Println(utc.UTC())
}
