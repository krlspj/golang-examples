package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // Use int64 for atomic operations
	var wg sync.WaitGroup

	// Start 10 goroutines that increment the counter atomically
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				// Increment the counter atomically
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
