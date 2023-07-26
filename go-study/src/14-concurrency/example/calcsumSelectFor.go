package main

import (
	"fmt"
	"time"
)

func evenSum2(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result
}
func squareSum2(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
	}
	ch <- result
}

func main() {
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum2(0, 100, evenCh)
	go squareSum2(0, 100, sqCh)

	//fmt.Println(<-evenCh + <-sqCh)

	// 只选择执行其中一个
	for {
		select {
		case x := <-evenCh:
			fmt.Println(x)
			return
		case y := <-sqCh:
			fmt.Println(y)
			return
		default:
			fmt.Println("Nothing available")
			time.Sleep(50 * time.Millisecond) // 阻止死锁
		}
	}
}

/*
Nothing available
2550
*/
