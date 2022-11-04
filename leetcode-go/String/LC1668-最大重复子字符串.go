package main

import "strings"

/**
 * 1668. 最大重复子字符串
 * 难度: 简单
 * 如果字符串 word 连续重复 k 次形成的字符串是 sequence 的一个子字符串，那么单词 word 的 重复值为 k。如果 word 不是 sequence 的子串，那么重复值 k 为 0 。 返回 最大重复值 k 。
 * 输入：sequence = "ababc", word = "ab"
 * 输出：2
 * 解释："abab" 是 "ababc" 的子字符串。
 * 输入：sequence = "ababc", word = "ac"
 * 输出：0
 * 解释："ac" 不是 "ababc" 的子字符串。
 */

func maxRepeating(sequence string, word string) int {
	count := 0
	for i := len(sequence) / len(word); i >= 1; i-- {
		if strings.Contains(sequence, strings.Repeat(word, i)) {
			return i
		}
	}
	return count
}

/*
执行结果：通过
执行用时：0 ms, 在所有 Go 提交中击败了 100% 的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了 24.24% 的用户

使用`strings.Repeat()`生成连续重复的子串
使用`strings.Contains()`判断子串是否存在
倒序遍历更快捷，最佳只需一次判断即可
*/
