package main

import "strings"

/**
 * BM83 字符串变形
 * 输入："This is a sample", 16
 * 输出："SAMPLE A IS tHIS"
 */

func trans(s string, n int) string {
	// write code here
	var res strings.Builder
	index := n
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			res.WriteString(s[i+1 : index])
			res.WriteByte(' ')
			index = i
		} else if i == 0 {
			res.WriteString(s[i:index])
		}
	}

	temp := res.String()
	// 字符大小写转换
	var ans strings.Builder
	for i := 0; i < len(temp); i++ {
		if temp[i] == ' ' {
			ans.WriteByte(' ')
		} else {
			ans.WriteByte(temp[i] ^ 32)
		}
	}

	return ans.String()
}

func main() {
	s := "This is a sample"
	println(trans(s, 16)) // SAMPLE A IS tHIS

}
