package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // Use int64 for atomic operations
	//var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&counter, 1)
				//counter.Add(1)

			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
