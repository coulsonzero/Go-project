package main

import (
	"fmt"
)

func out(start, end int, ch chan bool) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
	// 发送信息
	ch <- true
}

func main() {
	// 使用管道便不需要sleep，无需计算等待main()何时退出
	ch := make(chan bool)
	// 多线程异步方法，同时执行多个任务
	go out(0, 5, ch)
	go out(6, 100, ch)

	// 接受信息
	<-ch
	fmt.Println(<-ch)
}

/*
0
6
7
1
2
8
9
3
4
10
*/
