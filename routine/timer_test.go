package routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfterTimer(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())
	time := <-channel
	fmt.Println(time)
}

func TestAfterFuncTime(T *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	group.Wait()
	fmt.Println(time.Now())
}
