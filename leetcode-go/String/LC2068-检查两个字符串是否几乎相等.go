package main

import "strings"

/**
 * 执行结果：通过
 * 执行用时：0   ms, 在所有 Go 提交中击败了 100% 的用户
 * 内存消耗：1.9 MB, 在所有 Go 提交中击败了 100% 的用户
 */
func checkAlmostEquivalent(word1 string, word2 string) bool {
	var t int
	for _, v := range word1 {
		t = strings.Count(word1, string(v)) - strings.Count(word2, string(v))
		if t > 3 || t < -3 {
			return false
		}
	}

	for _, v := range word2 {
		t = strings.Count(word2, string(v)) - strings.Count(word1, string(v))
		if t > 3 || t < -3 {
			return false
		}
	}
	return true
}

/**
 * 执行结果：通过
 * 执行用时：4   ms, 在所有 Go 提交中击败了 12.82% 的用户
 * 内存消耗：1.9 MB, 在所有 Go 提交中击败了 33.33% 的用户
 */
func checkAlmostEquivalent2(word1 string, word2 string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(word1); i++ {
		m[word1[i]]++
		m[word2[i]]--
	}

	for _, v := range m {
		if v > 3 || v < -3 {
			return false
		}
	}
	return true
}

func main() {
	w1 := "zzzyyy"
	w2 := "iiiiii"
	println(checkAlmostEquivalent(w1, w2))
	println(checkAlmostEquivalent2(w1, w2))
}
