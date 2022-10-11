package main

/**
 * question: GP16 位运算
 * description: 已知a,b两个int类型变量，求出这两个变量的与，或，异或值，将结果依次存入切片中，然后返回这个切片。
 * input : 1,1
 * output: [1,1,0]
 */

func bitOperate(a int, b int) []int {
	// write code here
	res := []int{a & b, a | b, a ^ b}
	return res
}
