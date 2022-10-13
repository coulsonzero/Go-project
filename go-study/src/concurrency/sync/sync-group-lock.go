package main

import (
	"fmt"
	"sync"
)

func syncGroupUnlock() {
	var wg sync.WaitGroup
	var cnt int
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt++
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

// 未加锁
// 973
// 968
// ...

func syncGroupLock() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var cnt int
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			cnt++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

// 加锁
// 1000
// 1000

func syncGroupLock2() {
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
}

// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10

func main() {
	syncGroupLock2()
}
