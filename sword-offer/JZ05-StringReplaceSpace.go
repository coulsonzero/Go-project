package main

import (
	"bytes"
	"fmt"
	"strings"
)

/**
 * Q: JZ5-空字符替换
 * input : "We Are Happy"
 * output: "We%20Are%20Happy"
 */

func main() {
	s := "We Are Happy"
	// fmt.Println(replaceSpace(s))
	// fmt.Println(replaceSpace2(s))
	fmt.Println(replaceSpace3(s))
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
	// return strings.Replace(s, " ", "%20", -1)
}

func replaceSpace2(s string) string {
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

func replaceSpace3(s string) string {
	var buf bytes.Buffer
	// var buf strings.Builder
	for _, v := range []byte(s) {
		if string(v) == " " {
			buf.WriteString("%20")
		} else {
			buf.WriteByte(v)
		}
	}
	return buf.String()
}
