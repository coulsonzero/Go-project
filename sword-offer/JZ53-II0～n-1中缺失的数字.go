package main

/**
 * 输入: [0,1,3]
 * 输出: 2
 * 输入: [0,1,2,3,4,5,6,7,9]
 * 输出: 8
 */

func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}

// 二分查找
func missingNumber2(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == m {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}
