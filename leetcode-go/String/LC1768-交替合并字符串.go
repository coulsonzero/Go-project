package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	n := max(len(word1), len(word2))
	var res strings.Builder
	for i := 0; i < n; i++ {
		switch {
		case i >= len(word1):
			res.WriteByte(word2[i])
		case i >= len(word2):
			res.WriteByte(word1[i])
		default:
			res.WriteByte(word1[i])
			res.WriteByte(word2[i])
		}
	}
	return res.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(mergeAlternately("ab", "pqrs")) // apbqrs
	println(mergeAlternately("abcd", "pq")) // apbqcd
	println(mergeAlternately("abc", "pqr")) // apbqcr
}

/**
 * Question:    1768. 交替合并字符串
 * Description: 给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。
 * 如果一个字符串比另一个字符串长，就将多出来的字母追加到合后字符串的末尾。
 * 输入：word1 = "abc", word2 = "pqr"
 * 输出："apbqcr"
 * 解释：字符串合并情况如下所示：
 * word1：  a   b   c
 * word2：    p   q   r
 * 合并后：  a p b q c r
 *
 * 输入：word1 = "ab", word2 = "pqrs"
 * 输出："apbqrs"
 * 解释：注意，word2 比 word1 长，"rs" 需要追加到合并后字符串的末尾。
 * word1：  a   b
 * word2：    p   q   r   s
 * 合并后：  a p b q   r   s
 */

/*
 * +
 * 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
 * 内存消耗：2.4 MB, 在所有 Go 提交中击败了 7.83% 的用户
 */

/*
 * strings.Build
 * 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
 * 内存消耗：1.9 MB, 在所有 Go 提交中击败了 75.61% 的用户
 */

func mergeAlternately1(word1 string, word2 string) string {
	n := max(len(word1), len(word2))
	var res string
	for i := 0; i < n; i++ {
		if i >= len(word1) {
			res += string(word2[i])
		} else if i >= len(word2) {
			res += string(word1[i])
		} else {
			res += string(word1[i]) + string(word2[i])
		}
	}
	return res
}

func mergeAlternately2(word1 string, word2 string) string {
	n := max(len(word1), len(word2))
	var res strings.Builder
	for i := 0; i < n; i++ {
		if i >= len(word1) {
			res.WriteByte(word2[i])
		} else if i >= len(word2) {
			res.WriteByte(word1[i])
		} else {
			res.WriteByte(word1[i])
			res.WriteByte(word2[i])
		}
	}
	return res.String()
}

func mergeAlternately3(word1 string, word2 string) string {
	n := max(len(word1), len(word2))
	var res strings.Builder
	for i := 0; i < n; i++ {
		switch {
		case i >= len(word1):
			res.WriteByte(word2[i])
		case i >= len(word2):
			res.WriteByte(word1[i])
		default:
			res.WriteByte(word1[i])
			res.WriteByte(word2[i])
		}
	}
	return res.String()
}
