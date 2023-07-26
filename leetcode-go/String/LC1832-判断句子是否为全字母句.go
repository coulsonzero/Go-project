package main

import "strings"

func checkIfPangram(sentence string) bool {
	substr := "abcdefghijklmnopqrstuvwxyz"

	for _, v := range substr {
		if strings.Index(sentence, string(v)) == -1 {
			return false
		}
	}
	return true
}

// hashmap 哈希表法
func checkIfPangram(sentence string) bool {
	m := make(map[rune]bool)

	for _, v := range sentence {
		m[v] = true
	}

	return len(m) == 26
}

/*
class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        return len(set(sentence)) == 26
*/
