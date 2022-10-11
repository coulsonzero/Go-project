package main

import (
	"fmt"
	"sync"
)

// 多线程同步(等待n个线程完成后才执行下一步操作): sync.WaitGroup
// 信号量允许多个线程同时进入临界区

func main() {
	var wg sync.WaitGroup

	// 开启N个后台线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("hello %d \n", i)
			wg.Done()

		}(i)
	}

	// 等待N个后台线程完成
	wg.Wait()
	fmt.Println("end")
}

// hello 7
// hello 0
// hello 6
// hello 1
// hello 3
// hello 2
// hello 4
// hello 8
// hello 9
// hello 5
// end
