package main

import "sort"

/**
 * 输入：numbers = [3,4,5,1,2]
 * 输出：1
 * 输入：numbers = [2,2,2,0,1]
 * 输出：0
 */

// 数组比较
func minArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}

// 数组排序
func minArray2(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

// 二分查找
func minArray3(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[r] > nums[m] {
			r = m
		} else if nums[r] < nums[m] {
			l = m + 1
		} else {
			r--
		}
	}

	return nums[l]
}

/*
class Solution:
    def minArray(self, numbers: List[int]) -> int:
        return min(numbers)
*/
