package main

import (
	"fmt"
	"time"
)

// $ go build trace-debug.go
// $ GODEBUG=schedtrace=1000 ./trace-debug

// SCHED 0ms: gomaxprocs=8 idleprocs=7 threads=2 spinningthreads=0 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]
// SCHED 1004ms: gomaxprocs=8 idleprocs=8 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello World
// Hello World
// SCHED 2013ms: gomaxprocs=8 idleprocs=8 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello World
// SCHED 3022ms: gomaxprocs=8 idleprocs=8 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello World
// SCHED 4032ms: gomaxprocs=8 idleprocs=8 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello World

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}
