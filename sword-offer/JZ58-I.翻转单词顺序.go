package main

import (
	"strings"
)

/**
 * 注意：此题额外需要
 *      1.去除两端空格
 *      2.去除多余的连续空格，替换为单个空格
 */

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	// 反转字符串数组
	reverseSlice(arr)

	// 去除多余的空格
	temp := delDupSpace(arr)

	// 1.去除两端空格
	temp = strings.Trim(temp, " ")
	// 2.多个连续空格替换为单个空格
	return temp
}

func reverseSlice(arr []string) {
	r := len(arr) - 1
	for l := 0; l < r; l++ {
		arr[l], arr[r] = arr[r], arr[l]
		r--
	}
}

func delDupSpace(arr []string) string {
	// var res []string
	// for _, v := range arr {
	// 	if v != "" {
	// 		res = append(res, v)
	// 	}
	// }
	//
	// return strings.Join(res, " ")

	// 不占用额外数组空间
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return strings.Join(arr, " ")
}

func main() {
	println(reverseWords("a good   example"))
}

/*
class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join(s.split()[::-1])
*/
