package main

import (
	"fmt"
	"time"
)

func sum(start, end int, ch chan int) {
	res := 0
	for i := start; i <= end; i++ {
		// time.Sleep(50 * time.Millisecond)
		res += i
	}
	ch <- res
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// go实现并发,

	go sum(0, 5, ch1)
	go sum(6, 10, ch2)

	for {
		select {
		default:
			fmt.Println("receive default")
			// 避免死锁
			time.Sleep(50 * time.Millisecond)
		case x := <-ch1:
			fmt.Println("receive ch1")
			fmt.Println(x)
			return
		case y := <-ch2:
			fmt.Println("receive ch2")
			fmt.Println(y)
			return

		}
	}
}
