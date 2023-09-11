package main

//#include <math.h>
//#include <stdio.h>
//void printPiJSON() {
//	printf("from c: {\"pi\":%f}\n", M_PI);
//}
import "C"
import "fmt"

func main() {
	fmt.Printf("from go: {\"pi\":%f}\n", C.M_PI)
	C.printPiJSON()
}
