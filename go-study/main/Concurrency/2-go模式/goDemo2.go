package main

import (
	"fmt"
	"time"
)

func out2(start, end int) {
	for i := start; i <= end; i++ {
		// time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	// go实现并发,
	go out2(0, 5)
	go out2(6, 10)

	time.Sleep(500 * time.Millisecond)
}

/*
6
0
1
7
8
2
3
9
10
4
5
*/
