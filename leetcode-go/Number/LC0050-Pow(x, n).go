package main

/**
 * question: 50. Pow(x, n)
 * description: 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn )
 * 输入：x = 2.00000, n = 10
 * 输出：1024.00000
 * 输入：x = 2.10000, n = 3
 * 输出：9.26100
 * 输入：x = 2.00000, n = -2
 * 输出：0.25000
 * 解释：2-2 = 1/22 = 1/4 = 0.25
 */

func myPow(x float64, n int) float64 {
	var ans float64 = 1.0
	for i := n; i != 0; i /= 2 {
		if i%2 != 0 {
			ans *= x
		}
		x *= x
	}

	if n < 0 {
		ans = 1 / ans
	}
	return ans
}
