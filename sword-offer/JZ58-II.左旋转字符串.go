package main

/**
 * 输入: s = "abcdefg", k = 2
 * 输出: "cdefgab"
 * 输入: s = "lrloseumgh", k = 6
 * 输出: "umghlrlose"
 */

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
