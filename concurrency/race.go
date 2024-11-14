package main

import (
	"fmt"
	"sync"
)

// imaginemos una api donde ser hacen varias peticiones a servicios externos
// porejemplo 2 peticiones
// p1 -> A ...
// responde A -> p1 (t1 = 2s)
// p2 -> B ...
// responde B -> p2 (t2 = 3s)
// Cuanto tarda? :

// p1 -> A...
// p2 -> B...
// ?respA -> p1 (2s)
// ?respB -> p2 (3s)
// Cuanto tarda? :

func main() {
	var wg sync.WaitGroup
	var counter int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
	//fmt.Println("counter addr:", &counter)
}

// (M)(0xc11e088) -> (5)
// foo0 -> (M)(5) -> R -> foo0(5) -> (6) -> foo0: 6 -> W -> (M)(0xc11e088) -> (6)
// foo1 -> (M)(6) -> R -> foo0(6) -> (7) -> foo1: 7 -> W -> (M)(0xc11e088) -> (7)
// ... ...

// foo0 -> (5) -> R -> foo0(5) -> 6
// foo1 -> (5) -> R -> foo0(5)  -> 6
//-> foo0: 6 -> W -> (M)(0xc11e088) -> (6)
// -> foo1: 6 -> W -> (M)(0xc11e088) -> (6)
