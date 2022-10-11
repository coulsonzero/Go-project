package main

import (
	"fmt"
)

// 带缓存通道：需要将通道的发送方与接收方调换位置，否则无法

func main() {
	// bufferChan1()
	// bufferChan2()

}

func bufferChan1() {
	ch := make(chan int, 1) // 带缓存通道，size > 0
	go func() {
		fmt.Println("hello world !")
		// 发送数据
		ch <- 1
	}()

	// 接受数据
	<-ch
	// defer close(ch)
	fmt.Println("end")
}

// hello world !
// end

// 多线程同步
func bufferChan2() {
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
	fmt.Println("end")
}

// 随机顺序
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
// end

// func bufferChan3() {
// 	w := func() {
// 		fmt.Print("working...")
// 		time.Sleep(time.Second)
// 		fmt.Println("done")
// 	}
//
// 	ch := make(chan int, 5)
// 	works := 3
// 	for i := 0; i < works; i++ {
// 		go func() {
// 			ch <- 1
// 			w()
// 			<-ch
// 		}()
// 	}
// 	select {}
// }
