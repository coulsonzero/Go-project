package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var res []string
	var temp string
	for l := 0; l < len(s); l++ {
		for r := 0; r < len(s); r++ {
			if !strings.Contains(temp, string(s[r])) {
				temp += string(s[r])
			} else {
				fmt.Println(temp)
				res = append(res, temp)
				break
			}
		}
	}

	fmt.Println(res)
	return -1
}

func main() {
	s := "ababcbb"
	lengthOfLongestSubstring(s)
}
