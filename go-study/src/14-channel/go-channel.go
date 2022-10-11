package main

import "fmt"

func main() {
	ch := make(chan int)

	go say("coulson", ch)

	// 发送数据
	ch <- 1
}

func say(name string, ch chan int) {
	fmt.Printf("hi, %s\n", name)

	// 接收数据
	<-ch
}
