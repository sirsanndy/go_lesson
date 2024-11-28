package routine

import (
	"fmt"
	"go-routine/helper"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Test Channel"
		fmt.Println("End Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go helper.GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
