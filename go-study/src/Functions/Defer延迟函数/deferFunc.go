package main

import "fmt"

func main() {
	fmt.Println("start")

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i) // Stack结构: 先进后出
	// }

	fmt.Println("return", deferDemo())

	fmt.Println("end")
}

func deferDemo() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

/*
start
end
4
3
2
1
0
*/
