package main

import "fmt"

var msg string

func hello(ch chan bool) {
	msg = "hello world"
	ch <- true
}

func main() {
	ch := make(chan bool, 1)
	go hello(ch)
	<-ch
	fmt.Println(msg) // hello world
}
