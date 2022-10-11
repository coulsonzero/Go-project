package _6_TypeConver

import "strconv"

/**
 * Question: GP10 字符求和
 * Description: 给定两个字符串类型的数字，求这两个数字之和，并返回字符串形式
 * input: "12","34"
 * output: "46"
 */

func sum(a string, b string) string {
	// write code here
	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(b)
	return strconv.Itoa(num1 + num2)
}
