package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// go1.17新增 context简化对于处理单个请求的多个Goroutine协程之间与请求域的数据、超时和退出等操作，实现线程安全退出或超时的控制

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(time.Second)
	// cancel()

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	i := 1
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// fmt.Println("working")
			fmt.Printf("working %d seconds. \n", i)
		}
		i++
	}
}
