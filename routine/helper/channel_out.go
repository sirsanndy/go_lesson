package helper

import "fmt"

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
