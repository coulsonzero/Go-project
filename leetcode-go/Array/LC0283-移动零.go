package main

/**
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 */

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
}
