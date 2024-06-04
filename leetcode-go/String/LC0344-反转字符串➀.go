package main

import "fmt"

/**
 * 反转字符串
 * Input: s = ["h","e","l","l","o"]
 * Output: ["o","l","l","e","h"]
 * 要求: 原地修改, 空间复杂度O(1)
 * 难度: 简单
 */



func reverseString(s []byte) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

// 双指针
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}



// 前后原地反转
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))

	s2 := "leetcode"
	t := []byte(s2)
	reverseString(t)
	fmt.Println(string(t))

}
