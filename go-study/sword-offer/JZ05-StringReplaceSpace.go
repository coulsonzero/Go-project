package main

import (
	"fmt"
	"strings"
)

/**
 * Q: JZ5-空包字符替换
 * input : "We Are Happy"
 * output: "We%20Are%20Happy"
 */

func main() {
	s := "We Are Happy"
	fmt.Println(replaceSpace(s))
	fmt.Println(replaceSpace2(s))
}

func replaceSpace(s string) string {
	// write code here
	res := ""
	for _, c := range s {
		if c == ' ' {
			res += "%20"
		} else {
			res += string(c)
		}
	}
	return res
}

func replaceSpace2(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
	// return strings.Replace(s, " ", "%20", -1)
}
