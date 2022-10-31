package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 华为机考：字符串分割
// Input : "12abc-abCABc-4aB@"  3
// Output: "12abc-abc-ABC-4aB-@"

func main() {
	s := "12abc-abCABc-4aB@"
	k := 3
	// var s string
	// var k int
	// fmt.Print("please enter a string (s): ")
	// fmt.Scanln(&s)
	// fmt.Print("please enter a integer (k): ")
	// fmt.Scanln(&k)
	fmt.Printf("output: %s \n", convertString(s, k))
}

func convertString(s string, k int) string {
	var bu strings.Builder
	bu.WriteString(s[:strings.Index(s, "-")])
	// res, s := s[:strings.Index(s, "-")], s[strings.Index(s, "-"):]
	s = s[strings.Index(s, "-"):]
	var temp string
	var count int

	for i, v := range s {
		if string(v) != "-" {
			temp += string(v)
			count++
		}
		if count == k || i == len(s)-1 {
			// res += "-" + convert(temp)
			bu.WriteString("-")
			bu.WriteString(convert2(temp))
			count, temp = 0, ""
		}
	}
	// return res
	return bu.String()
}

func convert(temp string) string {
	// 统计 temp 大小写字符的数量
	m := make(map[string]int)
	for _, v := range temp {
		if unicode.IsLetter(v) {
			if string(v) >= "a" && string(v) <= "z" {
				m["lower"]++
			} else {
				m["upper"]++
			}
		}
	}
	// 如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；
	// 反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；
	// 大小写字母的数量相等时，不做转换
	if m["upper"] > m["lower"] {
		temp = strings.ToUpper(temp)
	} else if m["upper"] < m["lower"] {
		temp = strings.ToLower(temp)
	}
	return temp
}

func convert2(temp string) string {
	m := make(map[string]int)
	for _, v := range temp {
		if unicode.IsUpper(v) {
			m["upper"]++
		} else if unicode.IsLower(v) {
			m["lower"]++
		}
	}

	if m["upper"] > m["lower"] {
		temp = strings.ToUpper(temp)
	} else if m["upper"] < m["lower"] {
		temp = strings.ToLower(temp)
	}
	return temp
}
