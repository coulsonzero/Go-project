package main

import (
	"fmt"
	"strings"
)

/**
 * 输入：s = "We are happy."
 * 输出："We%20are%20happy."
 */


func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}


func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))	// We%20are%20happy.
}