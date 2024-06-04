package main

import "strings"

/**
 * Input:  "This is a sample",16
 * Output: "SAMPLE A IS tHIS"
 */

func trans( s string ,  n int ) string {
	// write code here
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		// 大小写转换
		if 'a' <= s[i] && s[i] <= 'z' {
			sb.WriteByte(s[i] - 32)
		} else if 'A' <= s[i] && s[i] <= 'Z' {
			sb.WriteByte(s[i] + 32)
		} else {
			sb.WriteByte(s[i])
		}
	}

	// 字符串反转
	arr := strings.Split(sb.String(), " ")
	i, j := 0, len(arr)-1
	for ; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, " ")
}