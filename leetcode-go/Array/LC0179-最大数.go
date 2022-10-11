package main

import (
	"sort"
	"strconv"
	"strings"
)

/**
 * 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
 * 输入：nums = [3,30,34,5,9]
 * 输出："9534330"
 */

func largestNumber(nums []int) string {
	// []int => []string
	var arr []string
	for _, num := range nums {
		arr = append(arr, strconv.Itoa(num))
	}

	// sort reverse string
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] >= arr[j]+arr[i]
	})

	// ps
	if arr[0] == "0" {
		return "0"
	}
	return strings.Join(arr, "")
}
