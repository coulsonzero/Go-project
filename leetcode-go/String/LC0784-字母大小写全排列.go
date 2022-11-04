package main

import (
	"fmt"
	"unicode"
)

/**
 * 784. 字母大小写全排列
 * 以 任意顺序 返回输出。
 * 输入：s = "a1b2"
 * 输出：["a1b2", "a1B2", "A1b2", "A1B2"]
 * 输入: s = "3z4"
 * 输出: ["3z4","3Z4"]
 */

func letterCasePermutation(s string) []string {
	var res []string
	var dsf func(arr []byte, i int)
	dsf = func(arr []byte, i int) {
		if i == len(arr) {
			res = append(res, string(arr))
			return
		}
		if unicode.IsDigit(rune(arr[i])) {
			dsf(arr, i+1)
			return
		}
		arr[i] = byte(unicode.ToUpper(rune(arr[i])))
		dsf(arr, i+1)
		arr[i] = byte(unicode.ToLower(rune(arr[i])))
		dsf(arr, i+1)
	}
	dsf([]byte(s), 0)
	return res
}

func main() {
	fmt.Printf("%v \n", letterCasePermutation("a1b2c3"))
}
