package main

/**
 * question: GP15 逻辑运算
 * description: 给定两个bool类型变量a,b，求出这两个bool类型变量的逻辑and，or，not a,not b的值，将依次存入一个切片中，然后返回这个切片。
 * input : true,true
 * output: [true,true,false,false]
 */

func logicalOperation(a bool, b bool) []bool {
	// write code here
	slice := []bool{a && b, a || b, !a, !b}
	return slice
}
