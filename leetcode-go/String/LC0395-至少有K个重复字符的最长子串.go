package main

import (
	"strings"
)

/**
 * 395. 至少有 K 个重复字符的最长子串
 * 难度：中等
 * 描述：给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
 * 输入：s = "ababbc", k = 2
 * 输出：5
 * 解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
 */

// 时间： 4  ms, 击败 32.91%
// 内存: 2.2 MB, 击败 21.37%
func longestSubstring(s string, k int) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	// fmt.Printf("%#v\n", m)
	for c := range m {
		if m[c] < k {
			n := 0
			// fmt.Printf("s: %s, %v\n", s, strings.Split(s, string(c)))
			for _, ss := range strings.Split(s, string(c)) {
				nn := longestSubstring(ss, k)
				if nn > n {
					n = nn
				}
			}
			return n
		}
	}

	return len(s)
}

func main() {
	println(longestSubstring("ababbc", 2))    // 5
	println(longestSubstring("ababbcaab", 3)) // 8
}

/*
// python3
// 时间：40 ms, 击败 67.93%
// 内存: 15.1 MB, 击败 48.83%

class Solution:
    def longestSubstring(self, s: str, k: int) -> int:
        if not s:
            return 0
        for c in set(s):
            if s.count(c) < k:
                return max(self.longestSubstring(t, k) for t in s.split(c))
        return len(s)
*/
