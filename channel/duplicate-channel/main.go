package main

import (
	"fmt"
	"time"
)

type myObj struct {
	in chan string
}

func main() {
	fmt.Println("vim-go")
	p := new(myObj)
	p.in = make(chan string, 10)

	go f1(p)
	go f2(p)

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		p.in <- fmt.Sprintf("main: %d", i)
	}

}

func f1(p *myObj) {
	for msg := range p.in {
		fmt.Println("f1", msg)
	}
}

func f2(p *myObj) {
	for msg := range p.in {
		fmt.Println("f2", msg)
	}
}
