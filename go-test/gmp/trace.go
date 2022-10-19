package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// GMP模型
// g--goroutine 协程
// m--thread    线程
// p--processor 处理器

// 1. go run trace.go
// 2. go tool trace trace.out
// 生成web可视化调度流程

func main() {
	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	fmt.Println("Hello World")
}
