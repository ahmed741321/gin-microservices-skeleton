package tasks

import (
	"log"
	"time"
)

// ExampleTask defines a background task to be executed by the worker.
func ExampleTask() {
	for {
		// Simulate task processing
		log.Println("Executing example task...")

		// Sleep for a specified duration before the next execution
		time.Sleep(10 * time.Second)
	}
}
