package main

import "strings"

func isPrefixString(s string, words []string) bool {
	var bf strings.Builder
	for _, word := range words {
		bf.WriteString(word)
		if bf.Len() >= len(s) {
			break
		}
	}
	return bf.String() == s
}

func isPrefixString2(s string, words []string) bool {
	for _, word := range words {
		if !strings.HasPrefix(s, word) {
			return false
		}
		s = s[len(word):]

		if len(s) == 0 {
			return true
		}
	}

	return strings.Join(words, "") == s
}

func main() {
	println(isPrefixString("iloveleetcode", []string{"i", "love", "leetcode", "apples"})) // true
	println(isPrefixString("iloveleetcode", []string{"apples", "i", "love", "leetcode"})) // false
	println(isPrefixString("ccccccccc", []string{"c", "cc"}))                             // false
	println(isPrefixString("a", []string{"aa", "aaaa", "banana"}))                        // false
}
