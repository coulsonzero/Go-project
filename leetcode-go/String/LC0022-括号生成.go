package main

import "strings"

func generateParenthesis(n int) []string {
	var ans []string
	var backtrack func(s []string, l, r int)
	backtrack = func(s []string, l, r int) {
		if len(s) == 2*n {
			ans = append(ans, strings.Join(s, ""))
			return
		}

		if l < n {
			s = append(s, "(")
			backtrack(s, l+1, r)
			s = s[:len(s)-1]
		}
		if r < l {
			s = append(s, ")")
			backtrack(s, l, r+1)
			s = s[:len(s)-1]
		}
	}
	backtrack(ans, 0, 0)
	return ans
}
