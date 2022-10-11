package main

import "fmt"

func test() {
	x := 8
	fmt.Println(x)
}

var y = 12 // y := 12 会报错, 只能用于局部变量

func main() {
	//fmt.Println(x) //  undefined: x
	fmt.Println(y) // 12
}
