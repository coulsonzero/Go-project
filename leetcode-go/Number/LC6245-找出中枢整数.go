package main

import "math"

/**
 * 6245. 找出中枢整数
 * 难度：简单
 * 给你一个正整数 n ，找出满足下述条件的 中枢整数 x ：
 * 1 和 x 之间的所有元素之和等于 x 和 n 之间所有元素之和。
 * 返回中枢整数 x 。如果不存在中枢整数，则返回 -1 。题目保证对于给定的输入，至多存在一个中枢整数。
 * 输入：n = 8
 * 输出：6
 * 解释：6 是中枢整数，因为 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21 。
 * 输入：n = 1
 * 输出：1
 * 解释：1 是中枢整数，因为 1 = 1 。
 * 输入：n = 4
 * 输出：-1
 * 解释：可以证明不存在满足题目要求的整数。
 */

// 时间复杂度：O(1); 空间复杂度：O(1)
func pivotInteger(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return x
	}
	return -1
}

func pivotInteger2(n int) int {
	sum_l := 0
	for i := 1; i <= n; i++ {
		sum_l += i
		sum_r := (i + n) * (n - i + 1) / 2
		if sum_l == sum_r {
			return i
		}
	}
	return -1
}

func pivotInteger3(n int) int {
	sum_l := 0
	sum_r := n * (n + 1) / 2
	for i := 1; i <= n; i++ {
		sum_l += i
		// fmt.Println(sum_l, sum_r)
		if sum_l == sum_r {
			return i
		}
		sum_r -= i
	}
	return -1
}
