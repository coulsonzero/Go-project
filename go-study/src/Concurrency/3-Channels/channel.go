package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go hello("couslon", ch)

	// 发送数据
	ch <- "Coulson"
	defer close(ch)
}

func hello(name string, ch chan string) {
	fmt.Printf("hello, %s \n", name)
	// 从通道接受数据
	<-ch
}
