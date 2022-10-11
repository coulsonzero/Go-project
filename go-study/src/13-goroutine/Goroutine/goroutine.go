package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * 实现线程同步的方法：
 * 1. 互斥量
 * 2. 缓存通道
 */

func main() {
	// 线程同步
	// syncLock()
	// channel()

	// 多线程并发
	// channelCache()

	// 并发线程等待同步
	// channelCacheConcurrent()
	// syncWait()
	// syncWait2()
	// maxConcurrency()
	syncWaitLock()
}

func syncWaitLock() {
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

func maxConcurrency() {
	ch := make(chan string, 32)
	searchByBing := func(s string) {
		fmt.Println("bing: ", s)

	}
	searchByGoogle := func(s string) {
		fmt.Println("google: ", s)
	}
	searchByBaidu := func(s string) {
		fmt.Println("baidu: ", s)
	}

	go func() {
		searchByBing("golang")
		ch <- "ok"
	}()
	go func() {
		searchByGoogle("golang")
		ch <- "ok"
	}()
	go func() {
		searchByBaidu("golang")
		ch <- "ok"
	}()
	fmt.Println(<-ch)
}

// 控制并发数(最大并发阻塞)
// func concurrentBlock() {
// 	limit := make(chan int, 5)
// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			limit <- 1
// 			fmt.Println("hello goroutine channel cache")
// 			<-limit
// 		}()
// 	}
// 	select {}
// }

// 并发线程等待
func syncWait() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("hello goroutine channel cache", i)
		}(i)
	}
	wg.Wait()
}

func syncWait2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", i)
		}(i)
	}
	wg.Wait()
}

// 带缓存通道并发
func channelCacheConcurrent() {
	done := make(chan int, 10)

	// 开启N个后台线程
	for i := 0; i < cap(done); i++ {
		go func(i int) {
			fmt.Println("hello goroutine channel cache", i)
			done <- 1
		}(i)
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

// 带缓存通道
func channelCache() {
	done := make(chan int, 2)

	go func() {
		fmt.Println("hello goroutine channel cache")
		done <- 1
	}()
	<-done

}

// 无缓存通道
func channel() {
	done := make(chan int)

	go func() {
		fmt.Println("hello goroutine")
		<-done
	}()

	done <- 1
}

// 互斥锁
func syncLock() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("hello sync")
		mu.Unlock()
	}()

	mu.Lock()
}
