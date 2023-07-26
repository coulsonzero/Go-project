package main

import (
	"sort"
	"strconv"
	"strings"
)

/**
 * 面试题45. 把数组排成最小的数
 * 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
 * 输入: [10,2]
 * 输出: "102"
 * 输入: [3,30,34,5,9]
 * 输出: "3033459"
 */

func minNumber(nums []int) string {
	var s []string
	for _, v := range nums {
		s = append(s, strconv.Itoa(v))
	}

	sort.Slice(s, func(a, b int) bool {
		return s[a]+s[b] < s[b]+s[a]
	})

	return strings.Join(s, "")
}
