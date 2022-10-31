package main

/**
 * Question: 可被三整除的偶数的平均值
 * Descript: 给你一个由正整数组成的整数数组 nums ，返回其中可被 3 整除的所有偶数的平均值。
 * 示例：
 * 输入：nums = [1,3,6,10,12,15]
 * 输出：9
 * 解释：6 和 12 是可以被 3 整除的偶数。(6 + 12) / 2 = 9 。
 * 输入：nums = [1,2,4,7,10]
 * 输出：0
 * 解释：不存在满足题目要求的整数，所以返回 0 。
 */

func averageValue(nums []int) int {
	var res int
	var count int
	for _, v := range nums {
		if v%3 == 0 && v%2 == 0 {
			res += v
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return res / count
}
