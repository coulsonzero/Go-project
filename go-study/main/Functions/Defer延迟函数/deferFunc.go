package main

import "fmt"

func main() {
	fmt.Println("start")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // Stack结构: 先进后出
	}

	fmt.Println("end")
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
