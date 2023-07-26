package main

import "github.com/coulsonzero/gopkg/pro/slices"

func main() {
	nums := []int{1, 3, 6, 2, 7}
	println(slices.Contains(nums, 5))
	println(slices.Contains(nums, 1))
}
