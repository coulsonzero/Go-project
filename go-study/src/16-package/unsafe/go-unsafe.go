package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var intSlice []int
	var strSlice []string

	// fmt.Printf("intSlice: %T, intSlice: %T \n", intSlice, strSlice) // intSlice: []int, intSlice: []string

	// fmt.Printf("%T \n", (*reflect.SliceHeader)(unsafe.Pointer(&intSlice)))
	// fmt.Printf("%T \n", (*[]int)(unsafe.Pointer(&intSlice)))

	pHdr := (*reflect.SliceHeader)(unsafe.Pointer(&intSlice))
	qHdr := (*reflect.SliceHeader)(unsafe.Pointer(&strSlice))

	pHdr.Data = qHdr.Data
	pHdr.Len = qHdr.Len
	pHdr.Cap = qHdr.Cap

	fmt.Printf("%T \n", pHdr)
	fmt.Printf("%T \n", qHdr)
}
