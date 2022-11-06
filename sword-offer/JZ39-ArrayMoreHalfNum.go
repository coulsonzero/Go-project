package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return -1
}

func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(nums))
}
