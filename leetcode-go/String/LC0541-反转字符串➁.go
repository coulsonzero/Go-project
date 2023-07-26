package main

/**
 * 541. 反转字符串 II
 * 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符
 * 输入：s = "abcdefg", k = 2
 * 输出："bacdfeg"
 */

func reverseStr(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		reverseString(t[i:min(i+k, len(s))])
	}
	return string(t)
}

func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
