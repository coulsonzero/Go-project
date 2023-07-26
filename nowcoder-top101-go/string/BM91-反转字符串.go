package main

/**
 * "abcd"
 * "dcba"
 */

func solve(s string) string {
	// write code here
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
