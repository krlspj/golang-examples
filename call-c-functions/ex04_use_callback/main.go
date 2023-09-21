package main

/*
#include <stdio.h>
extern void goCallback(int p0);

// Define a C function that takes a function pointer as an argument
static void call_cb(void (*callback)()) {
    printf("Calling the callback from C\n");
    callback(3); // Call the Go callback function from C
}
*/
import "C"
import (
	"fmt"
	"log"
)

// Define a Go callback function
//
//export goCallback
func goCallback(p0 C.int) {
	fmt.Println("Callback function called from Go")
	log.Println("log from go")
	fmt.Println("input number:", p0)
}

func main() {
	// Call the C function, passing the Go callback as an argument
	cFn := C.goCallback
	C.call_cb((*[0]byte)(cFn))
}

//package main
//
///*
//#include <stdio.h>
//
//// Declare the C callback function signature
//typedef void (*CallbackFunc)();
//
//extern void goCallback(void);
//
//
//// Define a C function that takes a callback function pointer and invokes it
//void callCallback(CallbackFunc callback) {
//    callback();
//}
//*/
//import "C"
//import (
//	"fmt"
//	"unsafe"
//)
//
//// Define a Go callback function
//// export goCallback
//func goCallback() {
//	fmt.Println("Callback function called from Go")
//}
//
//func main() {
//	// Call the C function, passing the Go callback as an argument
//	C.callCallback(C.CallbackFunc(unsafe.Pointer(C.goCallback)))
//}
