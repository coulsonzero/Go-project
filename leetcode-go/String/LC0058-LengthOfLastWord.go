package main

import "fmt"

/******************************************************************
 * Input: s = "   fly me   to   the moon  "                       *
 * Output: 4                                                      *
 * Explanation: The last word is "moon" with length 4.            *
 ******************************************************************/

// 反向遍历
func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for s[right] == ' ' {
		right--
	}
	cnt := 0
	for right >= 0 && s[right] != ' ' {
		right--
		cnt++
	}
	return cnt
}

func main() {
	s := "   fly me   to   the moon  "
	res := lengthOfLastWord(s)
	fmt.Println(res) // 4
}
