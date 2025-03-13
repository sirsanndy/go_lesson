package routine

import (
	"fmt"
	"testing"
	"time"
)

func worker(jobs <-chan int, results chan<- int, delayedFunc func()) {
	j := <-jobs // Receive a single job from the channel
	fmt.Println("worker processing job", j)
	time.Sleep(1 * time.Second)

	delayedFunc()

	results <- j * 2
	fmt.Println("Worker finished.")
}

func delayedTask() {
	fmt.Println("Delayed task executed!")
}

func TestDelayed(t *testing.T) {
	jobs := make(chan int, 1)    // Buffered channel of size 1 for one job
	results := make(chan int, 1) // Buffered channel of size 1 for one result

	// Start the worker goroutine
	go worker(jobs, results, delayedTask)

	// Send the job
	jobs <- 7   // Example job value
	close(jobs) // Close the channel after sending the job (important!)

	// Receive and print the result
	result := <-results
	fmt.Println("Result:", result)

	fmt.Println("Job processed and result received.")

	// Independently delayed function (demonstration)
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("This is an independently delayed function.")
	})
	time.Sleep(3 * time.Second) // Keep the main function running to see the output.
}
