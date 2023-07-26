package 二分查找_排序

func minNumberInRotateArray(nums []int) int {
	// write code here
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
