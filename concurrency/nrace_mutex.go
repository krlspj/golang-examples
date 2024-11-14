package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int64 // Use int64 for atomic operations
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
