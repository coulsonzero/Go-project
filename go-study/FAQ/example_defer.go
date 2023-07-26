package main

import (
	"fmt"
	"os"
)

func main() {
	println(example_defer3(1))
}

// defer: 多个defer遵循 栈"先进后出"顺序
// defer与return: defer在return之后执行，但在函数退出之前执行
func example_defer(num int) string {
	fmt.Printf("start: %d \n", num)
	defer fmt.Printf("defer: %d \n", num+1)
	fmt.Printf("middle: %d \n", num)
	defer fmt.Printf("defer2: %d \n", num+2)
	fmt.Printf("end: %d \n", num)
	return fmt.Sprintf("return: %d \n", num+3)
}

/*
start: 1
middle: 1
end: 1
defer2: 3
defer: 2
return: 4
*/

// defer 可以使用函数修改返回值
func example_defer2(num int) string {
	defer func() { num++; fmt.Printf("defer: %d \n", num) }()
	fmt.Printf("middle: %d \n", num)
	defer func() { num++; fmt.Printf("defer2: %d \n", num) }()
	fmt.Printf("end: %d \n", num)
	num = 5

	return fmt.Sprintf("return: %d \n", num)
}

/*
middle: 1
end: 1
defer2: 6
defer: 7
return: 5
*/

// defer： 函数中有os.Exit()退出程序时不再执行defer语句
func example_defer3(num int) string {
	defer func() { num++; fmt.Printf("defer: %d \n", num) }()
	fmt.Printf("middle: %d \n", num)
	defer func() { num++; fmt.Printf("defer2: %d \n", num) }()
	fmt.Printf("end: %d \n", num)
	num = 5
	os.Exit(1)
	return fmt.Sprintf("return: %d \n", num)
}

/*
middle: 1
end: 1
*/
