package main

import "fmt"

func main() {
	var (
		name string = "John"
		age  int    = 20
	)
	fmt.Printf("%s, %d", name, age) // John, 20

}

func test() {
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

func dataType() {
	var (
		name     string
		age      int
		gender   bool
		score    float64
		skills   []string
		scoreMap map[string]int
	)
	fmt.Printf("name =  %s, age = %d, gender = %t, score = %f \n", name, age, gender, score)
	// name = "", age = 0, gender = false, score = 0.000000
	fmt.Printf("%T, %T, %T, %T \n", name, age, gender, score)
	// string, int, bool, float64
	fmt.Printf("%v %T \n", skills, skills)
	// [] []string
	fmt.Printf("%v, %T \n", scoreMap, scoreMap)
	// map[], map[string]int
}
