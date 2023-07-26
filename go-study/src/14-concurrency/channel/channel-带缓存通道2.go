package main

// 多线程
func main() {
	ch := make(chan int, 10) // 带10个缓存

	for i := 0; i < cap(ch); i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	for i := 0; i < cap(ch); i++ {
		<-ch
	}
}

// hello 1
// hello 3
// hello 4
// hello 9
// hello 5
// hello 6
// hello 8
// hello 7
// hello 2
// hello 0
