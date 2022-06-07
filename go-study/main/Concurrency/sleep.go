package main

import (
	"fmt"
	"time"
)

func out(start, end int) {
	for i := start; i <= end; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	// 默认同步，先执行上行代码任务，在执行下个任务
	out(0, 5)
	out(6, 10)

	time.Sleep(50 * time.Millisecond)
}
