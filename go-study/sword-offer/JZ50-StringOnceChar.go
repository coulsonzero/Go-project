package main

import "fmt"

/**
 * Question: JZ50 第一个只出现一次的字符
 * description: 在一个长为 字符串中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数)
 * input : "google"
 * output: 4
 */

func FirstNotRepeatingChar(str string) int {
	// write code here
	cntMap := make(map[rune]int)
	for _, v := range str {
		cntMap[v]++
	}
	for i, v := range str {
		if cntMap[v] == 1 {
			return i
		}
	}
	return -1
	/*
	   cntMap := make(map[rune]int)
	   for _, v := range str {
	       cntMap[v]++
	   }
	   for i := 0; i < len(str); i++ {
	       if cntMap[[]rune(str)[i]] == 1 {
	           return i
	       }
	   }
	   return -1
	*/

	/*
	   cntMap := make(map[rune]int)
	   for k, v := range str {
	       if _, ok := cntMap[v]; !ok && strings.LastIndex(str, string(v)) == k {
	           return k
	       }
	       cntMap[v] = 1
	   }
	   return -1
	*/
}

func main() {
	str1 := "google"
	fmt.Println(FirstNotRepeatingChar(str1))

	str2 := "aac"
	fmt.Println(FirstNotRepeatingChar(str2))

	str3 := "zz"
	fmt.Println(FirstNotRepeatingChar(str3))
}
