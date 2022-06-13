package main

import (
	"fmt"
	"sort"
)

/**
 * Q: JZ3-数组中重复的数字
 * input : [2,3,1,0,2,5,3]
 * output: 2 or 3
 */

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(duplicate(nums))
}

func duplicate(nums []int) int {
	// write code here
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func duplicate2(nums []int) int {
	flag := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if flag[nums[i]] {
			return nums[i]
		} else {
			flag[nums[i]] = true
		}
	}

	return -1
}

func duplicate3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
