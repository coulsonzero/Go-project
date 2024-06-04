package main


/**
 * 给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。
 * 请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为1。
 * 输入：[3,4,5,1,2]
 * 输出：1
 * 输入：[2,2,2,0,1]
 * 输出：0
 */

func minArray(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l) >> 1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			r--
		}
	}
	return nums[l]
}