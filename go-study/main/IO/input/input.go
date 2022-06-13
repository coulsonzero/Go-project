package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 2000000000; i++ {

	}
	// end := time.Now()
	// secs := end.Sub(start)
	secs := time.Since(start).Seconds()
	// fmt.Println(secs)
	fmt.Printf("%.2fs", secs)

}

func input() {
	var name string
	fmt.Print("请输入名称(name): ")
	fmt.Scanln(&name)

	fmt.Println("Output: " + name)
}
