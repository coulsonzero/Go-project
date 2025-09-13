package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())                                // 2025-08-12 11:21:35.543884 +0800 CST m=+0.000155751
	fmt.Println(time.Now().Unix())                         // 1754968895
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))  // 2025-08-12 11:27:08
	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z")) // 2025-08-12T11:27:51Z
	fmt.Println(time.Now().Format(time.DateTime))          // 2025-08-12 11:27:51
}
