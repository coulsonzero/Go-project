package main

/**
 * 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
 * 输入：nums = [4,1,4,6]
 * 输出：[1,6] 或 [6,1]
 * 输入：nums = [1,2,10,4,1,4,3,3]
 * 输出：[2,10] 或 [10,2]
 */

func singleNumbers(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	var ans []int
	for k, v := range m {
		if v == 1 {
			ans = append(ans, k)
		}
	}

	return ans
}
