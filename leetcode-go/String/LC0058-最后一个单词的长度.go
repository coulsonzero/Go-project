package main

import (
	"fmt"
	"strings"
)

/**
 * 58. 最后一个单词的长度
 * 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度
 * Input: s = "   fly me   to   the moon  "
 * Output: 4
 * Explanation: The last word is "moon" with length 4.
 * 注意：需要去除字符串两端空格
 */

func lengthOfLastWord(s string) int {
	// 1.去除两端空格
	tmp := strings.Trim(s, " ")
	// 2.返回最后一个字符串的长度
	arr := strings.Split(tmp, " ")
	return len(arr[len(arr)-1])
}

func main() {
	s := "   fly me   to   the moon  "
	res := lengthOfLastWord(s)
	fmt.Println(res) // 4

}
