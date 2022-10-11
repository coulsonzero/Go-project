package main

/**
 * question: GP39 数字的阶乘
 * description: 一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，并且0的阶乘为1。自然数n的阶乘写作n!。
 * input : 3
 * output: 6  (1*2*3)
 */

func factorial(i int) int {
	// write code here
	if i == 0 || i == 1 {
		return 1
	}
	return i * factorial(i-1)
}
