package main

import "fmt"

// 无缓存通道
// ch <- ：发送数据
// <-ch  ：接收数据
// defer close(ch): 关闭通道

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello world !")
		// 从通道接受数据
		<-ch
	}()

	// 发送数据
	ch <- 12
	defer close(ch)
}
