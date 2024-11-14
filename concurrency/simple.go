package main

import (
	"fmt"
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
	var counter int
	//const max int = 1000 * 1000 * 1000
	const max int = 100

	for i := 0; i < 10; i++ {
		for j := 0; j < max; j++ {
			counter++
		}
	}

	fmt.Println("Final counter value:", counter)
}
