package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "4193 with words"
	fmt.Println(myAtoi(s)) // 4193

	s2 := "words and 987"
	fmt.Println(myAtoi(s2)) // 987

	s3 := "   -42"
	fmt.Println(myAtoi(s3)) // -42
}

func myAtoi(s string) int {
	m := map[rune]rune{
		'0': '0',
		'1': '1',
		'2': '2',
		'3': '3',
		'4': '4',
		'5': '5',
		'6': '6',
		'7': '7',
		'8': '8',
		'9': '9',
		'-': '-'}
	var res string
	for _, v := range s {
		if _, ok := m[v]; ok {
			res += string(v)
		}
	}
	ans, _ := strconv.Atoi(res)
	return ans
}
