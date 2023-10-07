package main

import (
	"fmt"
	"sync"
	"time"
)

var g_ch chan string

type myType struct {
	waiters sync.Map
}

func (t *myType) prepareChan(wg *sync.WaitGroup) {
	fmt.Println("------- prepare channel")

	key := "abc"
	waiter := make(chan string, 1)
	t.waiters.Store(key, waiter)
	defer close(waiter)
	defer t.waiters.Delete(key)

	fmt.Printf("channel waiter *p: %p\n", waiter)
	fmt.Println("------- send a to g_ch")
	g_ch <- key

	select {
	case resp := <-waiter:
		fmt.Println("----- response found", resp)
	case <-time.After(3 * time.Second):
		fmt.Println("timout found")
	}
	wg.Done()

}

func main() {
	g_ch = make(chan string)
	var aWg sync.WaitGroup

	aType := new(myType)
	aType.waiters = sync.Map{}

	go func() {
		fmt.Println("------- key:= <-g_ch")
		key := <-g_ch
		fmt.Println("------- key", key)

		w, ok := aType.waiters.Load(key)
		if ok {
			fmt.Printf("channel w *p: %p\n", w)
			w.(chan string) <- "message content"
		} else {
			fmt.Println("key not found")
		}
	}()

	aWg.Add(1)
	aType.prepareChan(&aWg)

	fmt.Println("--- print waiters sync.Map")
	aType.waiters.Range(func(key, value any) bool {
		fmt.Printf("--- key %s, val %v\n", key, value)
		return true
	})

	aWg.Wait()

	//	fmt.Println("waiting for 10 seconds")
	//	time.Sleep(10 * time.Second)

}
