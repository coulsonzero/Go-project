package main

/**
 * 1827. 最少操作使数组递增
 * 输入：nums = [1,1,1]
 * 输出：3
 * 解释：你可以进行如下操作：
 * 1) 增加 nums[2] ，数组变为 [1,1,2] 。
 * 2) 增加 nums[1] ，数组变为 [1,2,2] 。
 * 3) 增加 nums[2] ，数组变为 [1,2,3] 。
 * 输入：nums = [1,5,2,4,1]
 * 输出：14
 */

func minOperations(nums []int) int {
	ans := 0
	for i := 1; i < len(nums); i++ {
		ans += max(0, nums[i-1]+1-nums[i])
		nums[i] = max(nums[i], nums[i-1]+1)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
