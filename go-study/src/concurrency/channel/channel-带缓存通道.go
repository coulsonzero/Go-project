package main

import "fmt"

func main() {
	ch := make(chan int, 1) // 带缓存通道，带1个缓存
	go func() {
		fmt.Println("hello world !")
		ch <- 1 // 发送数据
	}()
	<-ch // 接受数据
	// defer close(ch)
}
