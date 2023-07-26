package main

import "sort"

func findMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

// 0ms
func findMin(nums []int) int {
	// 排序
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		// < 改为 <=: 因为有重复元素
		if nums[m] <= nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	return nums[l]
}

// 4ms
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] < nums[r] {
			r = m
		} else if nums[m] > nums[r] {
			l = m + 1
		} else {
			r--
		}
	}

	return nums[l]
}

func main() {
	println(findMin([]int{1, 3, 5}))       // 1
	println(findMin([]int{2, 2, 2, 0, 1})) // 0
	println(findMin([]int{3, 3, 1, 3}))    // 1
	println(findMin([]int{1, 3, 3}))       // 1
}

/*
class Solution:
    def findMin(self, nums):
        return min(nums)
*/
