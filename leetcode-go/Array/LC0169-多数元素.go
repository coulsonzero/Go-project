package main

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}

// hashmap 哈希表计数
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	n := len(nums)
	for _, v := range nums {
		m[v]++
		if m[v] > n/2 {
			return v
		}
	}
	return -1
}
