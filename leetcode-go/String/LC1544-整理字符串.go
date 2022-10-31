package main

import (
	"fmt"
	"strings"
)

// LC 1544. 整理字符串
// 给你一个由大小写英文字母组成的字符串 s 。
// 一个整理好的字符串中，两个相邻字符 s[i] 和 s[i+1] 要满足如下条件:
//
// 若 s[i] 是小写字符，则 s[i+1] 不可以是相同的大写字符。
// 若 s[i] 是大写字符，则 s[i+1] 不可以是相同的小写字符。
// 请你将字符串整理好，每次你都可以从字符串中选出满足上述条件的 两个相邻 字符并删除，直到字符串整理好为止。

// 递归法
func makeGood(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] && strings.ToLower(string(s[i])) == strings.ToLower(string(s[i+1])) {
			return makeGood(s[:i] + s[i+2:])
		}
	}
	return s
}

func main() {
	fmt.Println(makeGood("abBAcC"))     // ""
	fmt.Println(makeGood("abBAc"))      // "c"
	fmt.Println(makeGood("leEeetcode")) // "leetcode"
}
