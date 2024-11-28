package helper

import "time"

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Test Channel In"
}
