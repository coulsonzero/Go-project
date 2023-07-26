package main

import "strings"

/**
 * 输入：sentence = "leetcode exercises sound delightful"
 * 输出：true
 * 输入：sentence = "Leetcode is cool"
 * 输出：false
 */

// 直接遍历字符串判断
func isCircularSentence(sentence string) bool {
	n := len(sentence)
	for i := 0; i < n; i++ {
		if sentence[i] == ' ' && sentence[i+1] != sentence[i-1] {
			return false
		}
	}

	return sentence[0] == sentence[n-1]
}

// 字符串转数组判断
func isCircularSentence2(sentence string) bool {
	arr := strings.Split(sentence, " ")
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 单词的最后一个字符和下一个单词的第一个字符相等
		if arr[i][len(arr[i])-1] != arr[i+1][0] {
			return false
		}
	}
	// 最后一个单词的最后一个字符和第一个单词的第一个字符相等
	if arr[n-1][len(arr[n-1])-1] != arr[0][0] {
		return false
	}

	return true
}
