package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version()) // go1.18.3

	// num := 12
	// fmt.Println(toString(num))

	// slice := []int{1, 3, 5}
	// slice := []string{"apple", "banana", "orange"}
	// sliceToString(slice)

	fmt.Println(add[float64](4.2, 6.3))

}

func toString(num any) any {
	return num
}

func sliceToString(slice interface{}) {
	fmt.Printf("slice: %v \n", slice)
}

func sliceToString2(slice any) {
	fmt.Printf("slice: %v \n", slice)
}

func add[T ~float64 | ~int](a, b T) T {
	return a + b
}
