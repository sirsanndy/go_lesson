package routine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestMaxProc(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	total_cpu := runtime.NumCPU()
	fmt.Println("Total CPU", total_cpu)

	total_thread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", total_thread)
}

func TestRoutine(t *testing.T) {
	total_go_routine := runtime.NumGoroutine()
	fmt.Println("Total Go Routine", total_go_routine)
}
