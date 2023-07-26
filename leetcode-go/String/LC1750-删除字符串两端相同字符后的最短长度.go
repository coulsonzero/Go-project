package main

/**
 * 1750. 删除字符串两端相同字符后的最短长度
 * 难度：中等
 * 输入：s = "ca"
 * 输出：2
 * 输入：s = "cabaabac"
 * 输出：0
 * 输入：s = "aabccabba"
 * 输出：3
 * 输入："bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"
 * 输出：1
 * 输入："bbbbbbbbbbbbbbbbbbb"
 * 输出：0
 */

func minimumLength(s string) int {
	l := 0
	r := len(s) - 1
	for l < r && s[l] == s[r] {
		c := s[l]
		for l <= r && s[l] == c {
			l++
		}
		for l < r && s[r] == c {
			r--
		}
	}

	return r - l + 1
}
