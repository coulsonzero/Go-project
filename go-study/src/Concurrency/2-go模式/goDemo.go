package main

import (
	"fmt"
)

func out(from, to int) {
	for i := from; i <= to; i++ {
		fmt.Println(i)
	}
}

func main() {
	// 没有任何输出，因为main()于go并发前退出了，需要加sleep
	go out(0, 5)
	go out(6, 10)
}

// No Output
