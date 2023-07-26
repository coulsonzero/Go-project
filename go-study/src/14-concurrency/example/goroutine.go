package main

import (
	"fmt"
)

func out2(from, to int) {
	for i := from; i <= to; i++ {
		fmt.Println(i)
	}
}

func main() {
	// 没有任何输出，因为main()于go并发前退出了，需要加sleep
	go out2(0, 5)
	go out2(6, 10)
}

// No Output
