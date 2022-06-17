package main

import "C"
import (
	"fmt"
	"time"
)

/*
void Sum(int c) {
	int sum = 0;
    for (int i = 0; i <= c; i++) {
		sum += i;
	}
}
*/
// #cgo LDFLAGS:
import "C"

func main() {
	cycles := []int{500000, 1000000}
	counts := []int{10, 20, 100, 500}

	for _, v := range counts {
		for _, k := range cycles {
			// 使用Cgo(随着计算量的增加，CGo的性能优势逐渐体现，此时CGo的性能开销可以忽略不计了)
			timeStart := time.Now()
			for i := 0; i < k; i++ {
				C.Sum(C.int(v))
			}
			timeElapsed := time.Now().Sub(timeStart)
			fmt.Printf("count: %-3d, cycle: %-7d, cgo: %s, \n", v, k, timeElapsed)
			// 使用Go
			timeStartGo := time.Now()
			for i := 0; i < k; i++ {
				Sum(v)
			}
			timeElapsedGo := time.Now().Sub(timeStartGo)
			fmt.Printf("count: %-3d, cycle: %-7d, go: %s, \n", v, k, timeElapsedGo)
		}
	}
}

func Sum(c int) int {
	sum := 0
	for i := 0; i <= c; i++ {
		sum += i
	}
	return sum
}
