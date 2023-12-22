package main

import (
	"fmt"
	"time"
)

func waitForClose(ch <-chan string) {
	fmt.Println("loop")
	for {
		fmt.Println("select")
		select {
		case msg, b := <-ch:
			if b {
				fmt.Println("channel is open", msg)
			} else {
				fmt.Println("exit, channel is closed")
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func checkOpenChannel(ch chan<- string) {

	fmt.Println("loop-check")
	i := 0
	for {
		fmt.Println("select-check")
		msg := fmt.Sprintf("----- %d", i)
		select {
		case ch <- msg:
			fmt.Println("sended message:", msg)
		case <-time.After(10 * time.Second):
			fmt.Println("exit check open channel")
			return
		}
		i++
	}

}

func main() {
	buffChan1 := make(chan string, 1)
	go waitForClose(buffChan1)
	go checkOpenChannel(buffChan1)

	fmt.Println("waiting...")

	time.Sleep(11 * time.Second)

	close(buffChan1)
	time.Sleep(3 * time.Second)
	fmt.Println("exit program.")

}
