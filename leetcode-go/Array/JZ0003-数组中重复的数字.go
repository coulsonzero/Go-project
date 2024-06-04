package main

import (
	"fmt"
	"sort"
)

// map
func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] != 0 {
			return v
		}
		m[v]++
	}
	return -1
}

// sort array
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}


func main() {
	nums := []int{1, 3, 7, 13, 5, 3, 9}
	res := findRepeatNumber2(nums)
	fmt.Println(res)
}
