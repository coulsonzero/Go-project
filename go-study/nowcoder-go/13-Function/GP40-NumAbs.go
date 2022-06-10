package main

/**
 * question: GP40 绝对值
 * description: 定义一个函数，函数的功能是给出一个数，返回该数的绝对值。
 * input : -1
 * output: 1
 */

func absfunc(x int) int {
	// write code here
	if x < 0 {
		return -x
	}
	return x
}
