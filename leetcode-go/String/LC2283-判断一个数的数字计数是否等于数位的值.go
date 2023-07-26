package main

/**
 * 2283. 判断一个数的数字计数是否等于数位的值
 * 给你一个下标从 0 开始长度为 n 的字符串 num ，它只包含数字。
 * 如果对于 每个 0 <= i < n 的下标 i ，都满足数位 i 在 num 中出现了 num[i]次，那么请你返回 true ，否则返回 false 。
 * 输入：num = "1210"
 * 输出：true
 * 输入：num = "030"
 * 输出：false
 */

func digitCount(num string) bool {
	m := make(map[rune]int)
	for _, c := range num {
		m[c-'0']++
	}

	for i, c := range num {
		if m[rune(i)] != int(c-'0') {
			return false
		}
	}

	return true
}
