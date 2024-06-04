package main

import (
	"strings"
)

/**
 * 557. 反转字符串中的单词 III
 * 输入：s = "Let's take LeetCode contest"
 * 输出："s'teL ekat edoCteeL tsetnoc"
 */

// 空间复杂度: O(n)
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	for i, v := range arr {
		arr[i] = reverse_string([]byte(v))
	}
	return strings.Join(arr, " ")
}

func reverse_string(s []byte) string {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
	return string(s)
}

func main() {
	s := "Let's take LeetCode contest"
	println(reverseWords(s))		// s'teL ekat edoCteeL tsetnoc
}

/*
// python3

class Solution:
    def reverseWords(self, s: str) -> str:
        return ' '.join(i[::-1] for i in s.split())
*/
