package main

import "sort"

/**
 * 输入：arr = [3,2,1], k = 2
 * 输出：[1,2] 或者 [2,1]
 * 输入：arr = [0,1,2,1], k = 1
 * 输出：[0]
 */

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
