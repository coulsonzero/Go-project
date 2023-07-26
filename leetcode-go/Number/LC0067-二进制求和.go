package main

import (
	"math"
	"strconv"
)

func addBinary(a string, b string) string {
	math.Pow()
	n := max(len(a), len(b))
	carry := 0
	ans := ""
	for i := 0; i < n; i++ {
		if i < len(a) {
			carry += int(a[len(a)-1-i] - '0')
		}
		if i < len(b) {
			carry += int(b[len(b)-1-i] - '0')
		}

		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}

	if carry > 0 {
		ans = "1" + ans
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
