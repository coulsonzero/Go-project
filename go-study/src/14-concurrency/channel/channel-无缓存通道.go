package main

import (
	"fmt"
	"time"
)

// 线程同步：等待上一个操作完成后再执行下一个操作
// 当有一个线程在对内存进行操作时，其他线程都不可以对这个内存地址进行操作，直到该线程完成操作，
// 其他线程才能对该内存地址进行操作，而其他线程又处于等待状态

// channel: 通过通信来实现 共享内存，以及线程同步的

// 无缓存通道
// ch <- ：发送数据
// <-ch  ：接收数据
// defer close(ch): 关闭通道
// 如果带缓存(开启多个线程同步)则无法正确执行同步，需要调换接收方与发送方的位置

func main() {
	test()
}

func test() {
	ch := make(chan int) // 无缓存通道
	go func() {
		fmt.Println("hello world !")
		// 接受数据
		<-ch
	}()
	time.Sleep(time.Second * 3)
	// 发送数据
	ch <- 1
	// defer close(ch)

	// fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!

}
