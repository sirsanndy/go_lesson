package routine

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	group := &sync.WaitGroup{}
	data := &sync.Map{}
	var AddToMap = func(data *sync.Map, value int) {
		defer group.Done()
		group.Add(1)
		data.Store(value, value)
	}

	count := 1000
	for i := 0; i < count; i++ {
		go AddToMap(data, i)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
	fmt.Println("Selesai")
}
