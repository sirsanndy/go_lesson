package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data = list.New()
	data.PushBack("Muhamad")
	data.PushBack("Sandy")
	data.PushBack("Hasanudin")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Prev())
		fmt.Println(e.Value)
		fmt.Println(e.Next())
	}
}
