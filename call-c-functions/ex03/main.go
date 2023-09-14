package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lmylibrary
#include "mylibrary.h"
*/
import "C"
import "fmt"

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

func main() {
	a, b := 5, 3
	fmt.Println("a:", a, "b:", b)
	result0 := C.add(C.int(a), C.int(b))
	fmt.Println("result0:", result0)
	result := Add(a, b)
	fmt.Println("result:", result)
}
