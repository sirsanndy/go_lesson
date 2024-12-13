package routine

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func RunAsync(wait *sync.WaitGroup, counter int) {
	defer wait.Done()
	wait.Add(1)

	fmt.Println("Hello " + strconv.Itoa(counter))
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		go RunAsync(group, i)
	}

	group.Wait()
	fmt.Println("Selesai")
}
