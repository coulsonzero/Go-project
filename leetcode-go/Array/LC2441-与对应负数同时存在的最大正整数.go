package main

import "sort"

/**
 * 给你一个 不包含 任何零的整数数组 nums ，找出自身与对应的负数都在数组中存在的最大正整数 k 。
 * 返回正整数 k ，如果不存在这样的整数，返回 -1 。
 * 输入：nums = [-1,10,6,7,-7,1]
 * 输出：7
 * 解释：数组中存在 1 和 7 对应的负数，7 的值更大。
 */

func findMaxK(nums []int) int {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[j]+nums[i] == 0 {
			return nums[j]
		} else if nums[i]+nums[j] > 0 {
			j--
		} else {
			i++
		}
	}
	return -1
}
