package routine

import (
	"fmt"
	"go-routine/helper"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Test Channel 1"
		channel <- "Test Channel 2"
		channel <- "Test Channel 3"
	}()

	fmt.Println(<-channel)
	channel <- "Test Channel 4"
	fmt.Println(<-channel)
	channel <- "Test Channel 5"
	time.Sleep(3 * time.Second)
	fmt.Println("End of Channel")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- "Channel " + strconv.Itoa(i)
		}
	}()

	for data := range channel {
		fmt.Println("Received data", data)
	}

	fmt.Println("End of Channel")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go helper.SelectChannel(channel1, channel2)
	time.Sleep(3 * time.Second)
}
