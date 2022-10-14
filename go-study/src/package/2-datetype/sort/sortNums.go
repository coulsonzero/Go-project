package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 2, 6, 3, 1, 4}
	sortNums(nums)
	// reverseNums(nums)
	fmt.Println(nums)
}

func sortNums(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func sortNums2(nums []int) []int {
	sort.Sort(sort.IntSlice(nums))
	return nums
}

func reverseNums(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums
}
