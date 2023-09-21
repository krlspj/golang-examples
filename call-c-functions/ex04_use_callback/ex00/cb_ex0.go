package main

/*
typedef struct {
    void (*cb_f)();
} cb_s;

extern void cb_func(void);

static void cb_set(cb_s *s) {
    s->cb_f = &cb_func;
}
*/
import "C"

//export cb_func
func cb_func() {
	println("Callback function called from Go")
}

func main() {
	var x C.cb_s

	// Set the callback cb_f to point to the C function cb_func
	C.cb_set(&x)

	cbFunc := func() { C.cb_func() }
	// Call the function pointed to by cb_f
	cbFunc()
}
