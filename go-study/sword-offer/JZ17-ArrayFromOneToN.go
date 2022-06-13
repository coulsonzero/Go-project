package main

import "math"

/**
 * Question: JZ17 打印从1到最大的n位数
 * description: 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
 * input : 1
 * output: [1,2,3,4,5,6,7,8,9]
 *
 * tips: 1. 用返回一个整数列表来代替打印
 *       2. n 为正整数，0 < n <= 5
 */

func printNumbers(n int) []int {
	var res []int
	maxnum := int(math.Pow10(n))
	for i := 1; i < maxnum; i++ {
		res = append(res, i)
	}
	return res
}

func printNumbers2(n int) []int {
	// write code here
	var res []int
	// maxnum := int(math.Pow10(n))
	maxnum := 1
	for i := 0; i < n; i++ {
		maxnum = maxnum * 10
	}
	for i := 1; i < maxnum; i++ {
		res = append(res, i)
	}
	return res
}

func printNumbers3(n int) []int {
	// write code here
	m := map[int]int{
		1: 10,
		2: 100,
		3: 1000,
		4: 10000,
		5: 100000,
	}
	return printNum(m[n])
}

func printNum(length int) (res []int) {
	for i := 1; i < length; i++ {
		res = append(res, i)
	}
	return res
}
