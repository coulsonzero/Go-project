package main

import (
	"fmt"
	"strings"
)

/**
 * 实现 strStr()
 * 给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回-1 。
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 */

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))  // 2
	fmt.Println(strStr2(haystack, needle)) // 2
}

// strings.index()
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 双指针法
func strStr2(haystack string, needle string) int {
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}
