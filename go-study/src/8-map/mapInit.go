package main

import (
	"fmt"
	"sort"
)

func main() {

}


// 返回字符串中最多的字符
func mapCharCount() string {
	/*
		s := "yyds"
		m := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			if _, ok := m[s[i]]; ok {
				m[s[i]]++
			} else {
				m[s[i]] = 1
			}
		}
		fmt.Println(m)
		maxCount := 0
		var c byte
		for i, v := range m {
			if v > maxCount {
				maxCount = v
				c = i
			}
		}
		return string(c)
	*/
	s := "yyds"
	charArr := ([]byte)(s)
	charMap := make(map[byte]int)
	count := 0
	var ans byte
	for _, v := range charArr {
		charMap[v]++
		if charMap[v] > count {
			count = charMap[v]
			ans = v
		}
	}
	return string(ans)
}

func getNoRepeat() {
	s := []int{3, 3, 2, 2, 5, 5, 1, 2, 2}
	m := make(map[byte]int)
	for _, v := range s {
		m[byte(v)]++
	}
	// fmt.Println(m)
	res := []int{}
	for i, v := range m {
		if v == 1 {
			res = append(res, int(i))
		}
	}
	sort.Ints(res)
	fmt.Println(res)
}


