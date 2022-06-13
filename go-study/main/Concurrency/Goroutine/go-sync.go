package main

import (
	"fmt"
	"sync"
)

/**
 * 通过互斥量sync.Mutex来实现同步，互斥锁实现同步比较低级，可以改用无缓存通道来实现同步
 */

func main() {
	var mu sync.Mutex

	mu.Lock()
	go hi(&mu)
	mu.Lock()

}

func hi(mu *sync.Mutex) {
	fmt.Println("hello sync")
	mu.Unlock()
}
