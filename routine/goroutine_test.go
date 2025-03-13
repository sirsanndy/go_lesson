package routine

import (
	"fmt"
	"go-routine/helper"
	"testing"
	"time"
)

func TestCreateGoroutine(t *testing.T) {
	helper.RunHelloWorld()
	fmt.Println("Oops!")

	time.Sleep(1 * time.Second)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go helper.DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
