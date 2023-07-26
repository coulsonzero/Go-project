package main

import (
	"strings"
)

/*
输入：[奥] 卡夫卡
输出：卡夫卡

输入：惠特曼[美国][英]
输出：惠特曼

输入：[日] [村上春树
输出：[村上春树
*/

func main() {
	println(removeCountryInfo("[奥] 卡夫卡"))    // 卡夫卡
	println(removeCountryInfo("惠特曼[美国][英]")) // 惠特曼
	println(removeCountryInfo("[日] [村上春树"))  // [村上春树
}

func removeCountryInfo(s string) string {
	start := strings.Index(s, "[")
	end := strings.Index(s, "]")
	if start == -1 || end == -1 {
		return s
	}
	sub := s[start : end+1]
	s = strings.ReplaceAll(s, sub, "")
	s = strings.ReplaceAll(s, " ", "")

	return removeCountryInfo(s)
}
