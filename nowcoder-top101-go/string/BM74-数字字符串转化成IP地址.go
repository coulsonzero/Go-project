package main

import (
	"strconv"
	"strings"
)

var res []string

func restoreIpAddresses(s string) []string {
	// write code here
	res = nil
	dfs(s, nil, 0)
	return res
}

func dfs(s string, subStr []string, idx int) {
	if len(subStr) == 4 && idx == len(s) {
		res = append(res, strings.Join(subStr, "."))
		return
	}
	if len(subStr) == 4 && idx != len(s) {
		return
	}

	for i := 1; i <= 3; i++ {
		if i+idx > len(s) {
			return
		}

		if i != 1 && s[idx] == '0' {
			return
		}

		str := s[idx : idx+i]
		num, _ := strconv.Atoi(str)
		if i == 3 && num > 255 {
			return
		}

		dfs(s, append(subStr, str), idx+i)
	}
}
