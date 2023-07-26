package main

func findPeakElement(nums []int) int {
	index := 0
	for i, v := range nums {
		if v > nums[index] {
			index = i
		}
	}
	return index
}

func findPeakElement(nums []int) int {
	n := len(nums)
	for i := 1; i < n-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			return i
		}
	}

	if nums[0] > nums[n-1] {
		return 0

	}
	return n - 1
}

// 二分查找
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
