package main

import "strings"

func wordPattern(pattern string, s string) bool {
	m1 := make(map[string]byte)
	m2 := make(map[byte]string)
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i, word := range words {
		c := pattern[i]
		if m1[word] > 0 && m1[word] != c || m2[c] != "" && m2[c] != word {
			return false
		}
		m1[word] = c
		m2[c] = word
	}

	return true
}
