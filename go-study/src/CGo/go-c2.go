package main

/*
#include <stdio.h>

void hello(const char* name) {
	printf("hello, %s\n", name);
}

*/
import "C"

func main() {
	s := "coulson"
	cs := C.CString(s)
	C.hello(cs)
}
