package main

import "fmt"

func ping(ch1 chan<- string, msg string) {
	ch1 <- msg
}

func pong(ch1 <-chan string, ch2 chan<- string) {
	msg := <-ch1
	ch2 <- msg
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ping(ch1, "passed message") // ch1 <- "passed message"
	pong(ch1, ch2)              // ch2 <-  <-ch1
	fmt.Println(<-ch2)
}
