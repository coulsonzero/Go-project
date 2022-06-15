package main

import (
	"fmt"
)

func main() {

	fmt.Println("-- 1 --")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %s\n", r)
		}
		fmt.Println("-- 2 --")
	}()

	var slice = []int{1, 2, 3, 4, 5}

	slice[6] = 6
}

/*
-- 1 --
panic: runtime error: index out of range [6] with length 5
-- 2 --
*/
