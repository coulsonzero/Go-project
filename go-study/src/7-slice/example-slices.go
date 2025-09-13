package main

import "slices"

func main() {
	nums := []int{1, 3, 2, 7, 6}
	println(slices.Contains(nums, 5)) // false
	println(slices.Contains(nums, 2)) // true

}
