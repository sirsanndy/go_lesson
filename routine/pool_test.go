package routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	group := &sync.WaitGroup{}
	var pool = sync.Pool{
		//For default value
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Sandy")
	pool.Put("Hasanudin")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	group.Wait()
	fmt.Println("Selesai")
}
