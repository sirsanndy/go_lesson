package routine

import (
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var x int = 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	println("Counter = ", x)
}

func TestMutexRaceCondition(t *testing.T) {
	var mutex sync.Mutex
	var x int = 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	println("Counter = ", x)
}
