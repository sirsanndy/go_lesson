package routine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		time.Sleep(30 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}

	ticker.Stop()
	fmt.Println("Done")
}
