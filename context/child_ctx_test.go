package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestCtxWithValue(t *testing.T) {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")
	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")
	ctxF := context.WithValue(ctxC, "f", "F")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)

	fmt.Println(ctxF.Value("f"))
	fmt.Println(ctxF.Value("b"))
	fmt.Println(ctxF.Value("c"))
	fmt.Println(ctxF.Value("a"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Start Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("End", runtime.NumGoroutine())
}
