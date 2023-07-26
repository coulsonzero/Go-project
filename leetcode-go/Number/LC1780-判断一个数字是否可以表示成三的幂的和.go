package main

/**
 * 输入：n = 12
 * 输出：true
 * 解释：12 = 31 + 32
 * 输入：n = 91
 * 输出：true
 * 解释：91 = 30 + 32 + 34
 * 输入：n = 21
 * 输出：false
 */

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 > 1 {
			return false
		}
		n /= 3
	}
	return true
}
