package main

import "fmt"

// 多线程同步
func main() {
	ch := make(chan int, 10) // 带10个缓存

	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			fmt.Printf("hello %d \n", i)
			// 发送数据
			ch <- 1
		}(i)
	}

	for i := 0; i < cap(ch); i++ {
		<-ch
	}
}

// hello 1
// hello 3
// hello 4
// hello 9
// hello 5
// hello 6
// hello 8
// hello 7
// hello 2
// hello 0
