package main

import "fmt"

func main() {
	// 带缓存通道，带1个缓存
	ch := make(chan int, 1)
	go func() {
		fmt.Println("hello world !")
		// 发送数据
		ch <- 1
	}()
	// 接受数据
	<-ch

	defer close(ch)

	// 读取chan
	// a.已读取数据关闭chan后再读取会报错：fatal error: all goroutines are asleep - deadlock!
	// b.未读取关闭chan后读取可以读到数据，但再次读取会报错
	// fmt.Println(<-ch)

	// 写入chan
	// 带缓存通道chan关闭后再次写入chan一般不会报错，直到写入数量大于缓存大小
	// ch <- 2 // ok !
	// ch <- 3 // 超过缓存1：fatal error: all goroutines are asleep - deadlock!
}
