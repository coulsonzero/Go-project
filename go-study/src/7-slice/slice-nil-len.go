package main

import "fmt"

func main() {
	var slice1 []int
	fmt.Println(slice1 == nil)    // true
	fmt.Println(len(slice1) == 0) // true

	slice2 := []int{}
	fmt.Println(slice2 == nil)    // false
	fmt.Println(len(slice2) == 0) // true

	slice3 := make([]int, 0)      // 底层数组为空，但指针非空
	fmt.Println(slice3 == nil)    // false
	fmt.Println(len(slice3) == 0) // true
}
