package main

/**
 * 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
 * 输入: a = 1, b = 1
 * 输出: 2
 */

func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}

func add2(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}
