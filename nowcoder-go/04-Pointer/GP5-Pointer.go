package main

/**
 * Title: GP5 值和指针
 * Question: 给定两个变量a,b，判断两个变量的地址，值（a，b的地址取得）是否相等,将结果依次存入切片，并返回
 * input : 1,2
 * output: [false,false]
 */

func equal(a int, b int) []bool {
	// write code here
	return []bool{&a == &b, a == b}
}
