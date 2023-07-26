package main

/**
 * 输入：words = ["pay","attention","practice","attend"], pref = "at"
 * 输出：2
 * 解释：以 "at" 作为前缀的字符串有两个，分别是："attention" 和 "attend" 。
 */

import "strings"

func prefixCount(words []string, pref string) int {
	var cnt int
	for _, v := range words {
		if strings.HasPrefix(v, pref) {
			cnt++
		}
	}
	return cnt
}
