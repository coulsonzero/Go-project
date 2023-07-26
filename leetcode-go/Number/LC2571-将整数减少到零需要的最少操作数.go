package main

import "math/bits"

/*
输入：n = 39
输出：3
解释：我们可以执行下述操作：
- n 加上 20 = 1 ，得到 n = 40 。
- n 减去 23 = 8 ，得到 n = 32 。
- n 减去 25 = 32 ，得到 n = 0 。
可以证明使 n 等于 0 需要执行的最少操作数是 3

输入：n = 54
输出：3
解释：我们可以执行下述操作：
- n 加上 21 = 2 ，得到 n = 56 。
- n 加上 23 = 8 ，得到 n = 64 。
- n 减去 26 = 64 ，得到 n = 0 。
使 n 等于 0 需要执行的最少操作数是 3
*/

func minOperations(n int) int {
	return bits.OnesCount(uint(3*n ^ n))
}
