package main

import (
	"fmt"
	"sync"
)

/**
 * 互斥量实现并发
 */

func main() {
	syncWait()
}

func syncWait() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

/*
1
8
3
4
6
5
2
7
9
10
*/
