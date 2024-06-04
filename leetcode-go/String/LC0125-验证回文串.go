package main

import "strings"

/**
 * 验证回文串
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 输入: s = "A man, a plan, a canal: Panama"
 * 输出：true
 * 解释："amanaplanacanalpanama" 是回文串。
 * 难度: 简单
 */


func isPalindrome(s string) bool {
	str := ""
	for _, v := range s {
		// 只考虑字母和数字字符
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			// 忽略字母的大小写
			str += strings.ToLower(string(v))
		}
	}
	// 验证它是否是回文串
	return isReverseString(str)
}

func isReverseString(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func isReverseString2(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
