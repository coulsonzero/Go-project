package main

import "sort"

/**
 * question: 2475.数组中不等三元组的数目
 * description: nums[i]、nums[j] 和 nums[k] 两两不同
 * input : nums = [4,4,2,4,3]
 * output: 3
 * Explanation: 下面列出的三元组均满足题目条件：
 * (0, 2, 4) 因为 4 != 2 != 3
 * (1, 2, 4) 因为 4 != 2 != 3
 * (2, 3, 4) 因为 2 != 4 != 3
 * 共计 3 个三元组，返回 3
 */

/*
排序 + 分组统计:
时间复杂度：O(nlogN)
空间复杂度：O(1)
*/
func unequalTriplets(nums []int) int {
	sort.Ints(nums)
	ans, start := 0, 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			ans += start * (i - start + 1) * (n - 1 - i)
			start = i + 1
		}
	}
	return ans
}

// 时间复杂度：O(N)
// 空间复杂度：O(1)
func unequalTriplets2(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	a, c := 0, len(nums)
	for _, b := range cnt {
		c -= b
		ans += a * b * c
		a += b
	}
	return
}

// 暴力循环解法
// 时间复杂度：O(n^3)
func unequalTriplets1(nums []int) int {
	sort.Ints(nums)
	var ans, start int
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			ans += start * (i - start + 1) * (n - 1 - i)
			start = i + 1
		}
	}
	return ans
}
