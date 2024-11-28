package helper

import "time"

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Test Channel"
}
