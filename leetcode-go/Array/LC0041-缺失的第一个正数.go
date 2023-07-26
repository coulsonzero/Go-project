package main

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	t := 1
	for i := 0; i < len(nums); i++ {
		if t == nums[i] {
			t = nums[i] + 1
		}
	}

	return t
}
