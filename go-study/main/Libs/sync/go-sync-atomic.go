package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total uint64

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go Worker(&wg)
	go Worker(&wg)

	wg.Wait()

	fmt.Println(total) // 90
}

func Worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for ; i < 10; i++ {
		atomic.AddUint64(&total, i)
	}
}
