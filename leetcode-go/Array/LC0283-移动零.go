package main

/**
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 */

func moveZeroes(nums []int)  {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}




