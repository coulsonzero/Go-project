package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	Sum()
	end := time.Now()
	fmt.Println(end.Sub(start))

}

func Sum() int {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum += i
	}
	return sum
}
