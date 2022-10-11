package main

/**
 * 反转字符串
 * Input: s = ["h","e","l","l","o"]
 * Output: ["o","l","l","e","h"]
 * 要求: 原地修改, 空间复杂度O(1)
 * 难度: 简单
 */

// 字符串原地反转
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseString2(s []byte) {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
}
