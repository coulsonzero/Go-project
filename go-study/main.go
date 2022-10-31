package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())

	println(0 ^ 2)             // 2
	println(0 ^ 2 ^ 5)         // 7
	println(0 ^ 2 ^ 5 ^ 2)     // 5
	println(0 ^ 2 ^ 5 ^ 2 ^ 3) // 6
}
