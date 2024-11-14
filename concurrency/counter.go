package main

import (
	"fmt"
	"sync"
)

const routineNum int = 10
const countTo int = 1000 * 1000 * 1000

func countTo10000(wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	count := 0
	for i := 1; i <= countTo; i++ {
		count++
	}
	results <- count
}

func main() {
	var wg sync.WaitGroup
	results := make(chan int, routineNum) // 10 routines

	for i := 0; i < routineNum; i++ {
		wg.Add(1)
		go countTo10000(&wg, results)
	}

	wg.Wait()
	close(results)

	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Printf("Final sum: %d\n", sum)
}
