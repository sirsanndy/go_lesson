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
	time.Sleep(3 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go helper.GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(3 * time.Second)
}

func TestChannelIn(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go helper.OnlyIn(channel)
	time.Sleep(3 * time.Second)
}

func TestChannelOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go helper.OnlyOut(channel)
	time.Sleep(5 * time.Second)
}
