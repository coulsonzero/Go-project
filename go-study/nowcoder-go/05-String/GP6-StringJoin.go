package main

import "strings"

/**
 * Question: GP6 拼接字符串
 * Descript: 给定一个字符串数组，将其拼接成一个字符串
 * input :   ["hello","-","world"]
 * output:   "hello-world"
 */

func join(s []string) string {
	// write code here
	res := ""
	for _, v := range s {
		res += v
	}
	return res
}

func join2(s []string) string {
	return strings.Join(s, "")
}
