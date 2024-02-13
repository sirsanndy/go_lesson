package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Muhamad")
	data.PushBack("Hasanudin")
	var tail *list.Element = data.Back()
	data.InsertBefore("Sandy", tail)

	for txt := data.Front(); txt != nil; txt = txt.Next() {
		fmt.Println(txt.Value)
	}
}
