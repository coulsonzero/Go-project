package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("hello world !")
	}()

	go func() {
		fmt.Println("hello runtime !")
	}()

	goroutine := runtime.NumGoroutine()
	fmt.Printf("%v, %T", goroutine, goroutine)
}

func getReadMemStats() {
	var rm runtime.MemStats
	runtime.ReadMemStats(&rm)
	fmt.Printf("%+v\n", rm)
	fmt.Printf("alloc: %d\n", rm.Alloc)
	fmt.Printf("os: %d\n", rm.Sys)
}
