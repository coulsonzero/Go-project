package main

import "sort"

/**
 * @title：1760. 袋子里最少数目的球
 * @descript：给你一个整数数组 nums ，其中 nums[i] 表示第 i 个袋子里球的数目。同时给你一个整数 maxOperations 。
 * 			  你可以进行如下操作至多 maxOperations 次：
 * 			  你的开销是单个袋子里球数目的 最大值 ，你想要 最小化 开销。请你返回进行上述操作后的最小开销。
 * 输入：nums = [9], maxOperations = 2
 * 输出：3
 * 解释：
 * 	    将装有 9 个球的袋子分成装有 6 个和 3 个球的袋子。[9] -> [6,3] 。
 * 		将装有 6 个球的袋子分成装有 3 个和 3 个球的袋子。[6,3] -> [3,3,3] 。
 * 		装有最多球的袋子里装有 3 个球，所以开销为 3 并返回 3 。
 */

func minimumSize(nums []int, maxOperations int) int {
	sort.Ints(nums)
	l := 1
	r := nums[len(nums)-1]
	for l < r {
		mid := l + (r-l)/2

		ans := 0
		for _, v := range nums {
			ans += (v - 1) / mid
		}

		if ans <= maxOperations {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
