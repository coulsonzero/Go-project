package main

/**
 * 229. 多数元素 II
 * 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
 * 输入：nums = [3,2,3]
 * 输出：[3]
 */

func majorityElement(nums []int) []int {
	var ans []int
	m := make(map[int]int)
	n := len(nums)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v > n/3 {
			ans = append(ans, k)
		}
	}

	return ans
}
