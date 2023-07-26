package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(&wg, &mu)
	}
	wg.Wait()
}

func work(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	// mu.Lock()
	count++
	fmt.Println(count)
	// mu.Unlock()
}
