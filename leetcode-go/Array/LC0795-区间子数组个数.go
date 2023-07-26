package main

/**
 * question: 795. 区间子数组个数
 * 难度：中等
 * descript: 给你一个整数数组 nums 和两个整数：left 及 right 。
 *           找出 nums 中连续、非空且其中最大元素在范围 [left, right] 内的子数组，并返回满足条件的子数组的个数。
 * 输入：nums = [2,1,4,3], left = 2, right = 3
 * 输出：3
 * 解释：满足条件的三个子数组：[2], [2, 1], [3]
 */

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans := 0
	p0, p1 := -1, -1
	for i, v := range nums {
		if v > right {
			p0 = i
		}
		if v >= left {
			p1 = i
		}
		ans += p1 - p0
	}
	return ans
}
