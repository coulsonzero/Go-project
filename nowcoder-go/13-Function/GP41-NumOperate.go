package main

/**
 * question: GP41 加减乘除
 * description: 定义一个函数，实现输入a，b两个数，返回两数的和，差，乘积，商，余数。然后依次存入切片中，最后返回。
 * input : 2,2
 * output: [4,0,4,1,0]
 */

func operate(a int, b int) []int {
	// write code here
	return []int{a + b, a - b, a * b, a / b, a % b}
}
