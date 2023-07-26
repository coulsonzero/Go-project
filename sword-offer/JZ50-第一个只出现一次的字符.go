package main

import "strings"

// 时间复杂度：O(N), 44ms
func firstUniqChar(s string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++

	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}

// 时间复杂度：O(N^2), 160ms
func firstUniqChar2(s string) byte {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) == 1 {
			return s[i]
		}
	}
	return ' '
}
