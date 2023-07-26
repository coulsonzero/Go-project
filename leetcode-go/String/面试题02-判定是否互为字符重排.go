package main

import (
	"reflect"
)

/**
 * 输入: s1 = "abc", s2 = "bca"
 * 输出: true
 * 输入: s1 = "abc", s2 = "bad"
 * 输出: false
 */

func CheckPermutation(s1 string, s2 string) bool {
	m1 := make(map[rune]int)
	for _, v := range s1 {
		m1[v]++
	}

	m2 := make(map[rune]int)
	for _, v := range s2 {
		m2[v]++
	}

	return reflect.DeepEqual(m1, m2)
}
