package main

import "fmt"

func main() {
	// 1. 变量声明：未初始化时有默认值
	var a int
	// 2. 变量声明并初始化
	var pi float64 = 3.14
	var pi2 = 3.14
	// 3. 变量声明并初始化
	num := 8

	fmt.Printf("value = %d, type of %T\n", a, a)     // a = 0, type of int
	fmt.Printf("value = %f, type of %T\n", pi, pi)   // value = 3.140000, type of float64
	fmt.Printf("value = %f, type of %T\n", pi2, pi2) // value = 3.140000, type of float64
	fmt.Printf("value = %d, type of %T\n", num, num)

	// 声明多个变量

	var (
		name string = "john"
		age  int    = 20
	)
	// var name, age = "john", 20
	fmt.Println("name: ", name, "age: ", age)

	x, y := "john", 20
	fmt.Println("name: ", x, "age: ", y)
}
