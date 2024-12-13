package routine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOne() {
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			once.Do(OnlyOne)
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
