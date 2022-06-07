package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(600 * time.Millisecond)
	end := time.Now()
	fmt.Println(end.Sub(start)) // 606.114625ms

}
