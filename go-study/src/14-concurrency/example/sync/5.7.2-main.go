package main

import "go-study/src/concurrency/example/sync/pkg"

func main() {
	var ch = make(chan string, 2)
	// 结束标志
	var finished = make(chan bool)

	go pkg.TaskProducer(ch)
	go pkg.TaskConsumer(ch, finished)

	<-finished
}

/*
Task2:9,7
Task1:1,7
Task5:0,5
Task8:0,1
Task6:1,8
Task7:4,2
Task4:6,9
Task3:8,4

*/
