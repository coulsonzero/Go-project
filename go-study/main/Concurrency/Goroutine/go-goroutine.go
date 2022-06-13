package main

import "fmt"

/**
 * main()退出时程序结束，不会等待任何后台线程
 * 可以通过同步原语来给两个事件明确顺序，实现同步
 * 无缓存通道实现同步
 */

func main() {
	done := make(chan int)

	go hello(done)

	<-done // select{}、for{}等方法可以阻塞main()线程，从而避免程序过早退出; 如果需要正常退出, 可以调用os.Exit(0)实现。
}

func hello(done chan int) {
	fmt.Println("hello goroutine")

	done <- 1 // close(done) 亦可

}
