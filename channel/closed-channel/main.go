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

func (t *myType) prepareChan() {
	fmt.Println("------- prepare channel")
	waiter := make(chan string, 1)
	t.waiters.Store("a", waiter)
	fmt.Println("------- send a to g_ch")
	g_ch <- "a"

	select {
	case resp := <-waiter:
		fmt.Println("----- response found", resp)
	case <-time.After(3 * time.Second):
		fmt.Println("timout found")
	}

}

func main() {
	g_ch = make(chan string)

	aType := new(myType)
	aType.waiters = sync.Map{}

	go func() {
		fmt.Println("------- key:= <-g_ch")
		key := <-g_ch
		fmt.Println("------- key", key)

		w, ok := aType.waiters.Load(key)
		if ok {
			w.(chan string) <- "message content"
		} else {
			fmt.Println("key not found")
		}
	}()

	go aType.prepareChan()

	//	fmt.Println("waiting for 10 seconds")
	//	time.Sleep(10 * time.Second)

}
