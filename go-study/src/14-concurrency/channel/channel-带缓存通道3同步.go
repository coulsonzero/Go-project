package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	ch := make(chan int, 10) // 带10个缓存
	cnt := 0
	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			defer func() { ch <- 1 }()
			mu.Lock()
			cnt++
			fmt.Println("cnt: ", cnt)
			mu.Unlock()
		}(i)
	}

	for i := 0; i < cap(ch); i++ {
		<-ch
	}
}

/*
cnt:  1
cnt:  2
cnt:  3
cnt:  4
cnt:  5
cnt:  6
cnt:  7
cnt:  8
cnt:  9
cnt:  10
*/
