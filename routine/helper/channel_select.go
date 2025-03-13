package helper

import "fmt"

func SelectChannel(channel1 chan string, channel2 chan string) {
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	select {
	case data := <-channel1:
		fmt.Println("Data dari Channel 1", data)
		break
	case data := <-channel2:
		fmt.Println("Data dari Channel 2", data)
		break
	}

	select {
	case data := <-channel1:
		fmt.Println("Data dari Channel 1", data)
		break
	case data := <-channel2:
		fmt.Println("Data dari Channel 2", data)
		break
	}
}
