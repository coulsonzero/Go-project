package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	select4()
}

func select1() {
	ch := make(chan int)
	go func() {
		select {
		case ch <- 0:
			fmt.Println("0")
		case ch <- 1:
			fmt.Println("1")
		default:
			fmt.Println("default")
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// fatal error: all goroutines are asleep - deadlock!

func select2() {
	ch := make(chan bool)
	go work(ch)

	time.Sleep(time.Second)
	ch <- true
}

// working
// ... long time

func work(ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("end")
		default:
			fmt.Println("working")
		}
	}
}

func select3() {
	ch := make(chan bool)
	for i := 0; i < 5; i++ {
		go work(ch)
	}

	time.Sleep(time.Second)
	close(ch)
}

// working
// ... long time

func select4() {
	ch := make(chan bool)

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(ch, &wg)
	}

	time.Sleep(time.Second)
	close(ch)
	wg.Wait()
}

func worker(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Println("end")
		default:
			fmt.Println("working")
		}
	}
}
