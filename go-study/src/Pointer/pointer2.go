package main

import "fmt"

func set(val int) {
	val = 8
}

func set_ptr(ptr *int) {
	*ptr = 8
}

func main() {
	x := 12

	set(x)
	fmt.Println(x) // 12

	set_ptr(&x)
	fmt.Println(x) // 8
}
