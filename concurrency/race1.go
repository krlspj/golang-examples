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
		go func(i *int, wg *sync.WaitGroup) {
			defer (*wg).Done()
			for j := 0; j < 1000; j++ {
				(*i)++
			}
		}(&counter, &wg)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
