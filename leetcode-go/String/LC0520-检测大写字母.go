package main

/**
 * 我们定义，在以下情况时，单词的大写用法是正确的：
 * 全部字母都是大写，比如 "USA" 。
 * 单词中所有字母都不是大写，比如 "leetcode" 。
 * 如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
 * 给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。
 * 输入：word = "USA"
 * 输出：true
 * 输入：word = "FlaG"
 * 输出：false
 * 输入："Leetcode"
 * 输出：true
 */

func detectCapitalUse(word string) bool {
	upper := 0
	lower := 0
	n := len(word)
	for _, c := range word {
		if 'a' <= c && c <= 'z' {
			lower++
		} else if 'A' <= c && c <= 'Z' {
			upper++
		}
	}

	// fmt.Printf("upper: %d, lower: %d \n", upper, lower)
	return upper == n || lower == n || ('A' <= word[0] && word[0] <= 'Z' && upper == 1)
}
