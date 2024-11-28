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
	chan_in := make(chan string)
	defer close(chan_in)
	go helper.OnlyIn(chan_in)
	fmt.Println(<-chan_in)
	time.Sleep(3 * time.Second)
}

func TestChannelOut(t *testing.T) {
	chan_out := make(chan string)
	defer close(chan_out)

	go helper.OnlyOut(chan_out)
	time.Sleep(3 * time.Second)
}
