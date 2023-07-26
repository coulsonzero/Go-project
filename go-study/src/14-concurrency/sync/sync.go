package main

import (
	"fmt"
	"sync"
)

// 互斥量：sync.Mutex实现线程同步
// 互斥锁只允许一个线程进入临界区

// 传统线程并发中限制共享资源访问，确保了在当前运行环境下有且只有一个线程对共享资源进行访问，防止出现资源竞争
// 有一个苹果，但是有两个人都想拿这个苹果，到底该给谁呢，这就是资源竞争，
// 拥有锁的人才可以得到苹果而另一个人只能等拥有锁的人解锁的时候才能拿这个苹果。

func main() {
	var mu sync.Mutex // {0, 0}
	// 锁定互斥量
	mu.Lock() // {1, 0}
	go func() {
		fmt.Println("hello world !")
		// 解锁互斥量
		mu.Unlock() // {2, 1}
	}()
	mu.Lock() // {1, 0}

	fmt.Println("end")
}

// hello world !
// end
