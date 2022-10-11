package main

import "fmt"

/**
 * Question: GP8 回文数
 * Descript: 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 *           回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数
 * input : 121
 * output: true
 */

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
