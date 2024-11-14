package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

//
//func main() {
//	var wg sync.WaitGroup
//	var val int = 0
//	var counter *int = &val
//
//	// Start 10 goroutines that all increment the counter
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(i *int) {
//			defer wg.Done()
//			for j := 0; j < 10000; j++ {
//				// Data race: multiple goroutines accessing the shared variable 'counter'
//				(*i)++
//			}
//		}(counter)
//	}
//
//	// Wait for all goroutines to finish
//	wg.Wait()
//	fmt.Println("Final counter value:", *counter)
//}
