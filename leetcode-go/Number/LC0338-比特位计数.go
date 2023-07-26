package main

/**
 * question: 338. 比特位计数
 * description: 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，
 *              计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
 * 输入：n = 2
 * 输出：[0,1,1]
 * 解释：
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 *
 * 输入：n = 5
 * 输出：[0,1,1,2,1,2]
 * 解释：
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 * 3 --> 11
 * 4 --> 100
 * 5 --> 101
 */

// 0ms 计算二进制中1的个数
func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = hammingWeight(i)
	}

	return ans
}

// 二进制中1的个数
func hammingWeight(num int) int {
	var res int
	for n := num; n != 0; n /= 2 {
		res += n % 2
	}
	return res
}

// 8ms
func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = strings.Count(strconv.FormatInt(int64(i), 2), "1")
	}

	return ans
}

// 4ms 位运算
func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}

	return ans
}

// 4ms
func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}

	return ans
}
