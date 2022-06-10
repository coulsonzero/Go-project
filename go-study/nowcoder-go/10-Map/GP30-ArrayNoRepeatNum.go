package main

import "sort"

/**
 * question: GP30 不重复的数
 * description: 给定一个数组，找出数组中所有不重复的数字，并按照从小到大的顺序进行输出
 * input : [3,3,2,2,5,5,1,2,2]
 * output: [1]
 */
func getNoRepeat(s []int) []int {
	// write code here
	m := make(map[byte]int)
	for _, v := range s {
		m[byte(v)]++
	}
	res := []int{}
	for i, v := range m {
		if v == 1 {
			res = append(res, int(i))
		}
	}
	sort.Ints(res)
	return res
}
