package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	nums = insert(nums, 1, 6, 4)
	fmt.Printf("res: %v\n", nums) 		// res: [1 6 4 2 3]
}

type sl interface {
	int | int64 | float64 | string | bool
}

func insert(slice []int, index int, value ...int) []int {
	return append(slice[:index], append(value, slice[index:]...)...)
}



