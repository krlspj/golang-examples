package main

import (
	"fmt"
	"sync"
)

func main() {
	aMap := sync.Map{}

	aMap.Store("k0", "v0")
	fmt.Println("the map:", &aMap)
	var mlen int
	aMap.Range(func(key, value any) bool {
		mlen++
		return true
	})
	fmt.Println("the map len:", mlen)
	aMap.Store("k1", "v1")
	aMap.Store("k2", "v2")
	aMap.Store("k3", "v3")
	//aMap.Store("k1", "v1")
	mapStats(aMap)

	resp, ok := aMap.Load("k0")
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println("resp", resp.(string))
	}

	mapStats(aMap)
	aMap.Delete("k0")
	mapStats(aMap)

	cleanMap(&aMap)
	mapStats(aMap)

}

func mapStats(m sync.Map) {
	fmt.Println("the map:", m)
	var mlen int
	m.Range(func(key, value any) bool {
		mlen++
		return true
	})
	fmt.Println("the map len:", mlen)
}

func cleanMap(m *sync.Map) {
	m.Range(func(key, value any) bool {
		m.Delete(key)
		return true
	})
}
