package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			fmt.Println(count)
			mu.Unlock()
		}()
	}
	wg.Wait()

	// fmt.Println(count)
}

/*
1
2
3
4
5
6
7
8
9
10
*/
