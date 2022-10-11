package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "loveleetcode"
	s := "aabb"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) bool {
	for i := 0; i < len(s); i++ {
		if strings.Index(s, string(s[i])) != strings.LastIndex(s, string(s[i])) {
			return false
		}
	}
	return true
}
