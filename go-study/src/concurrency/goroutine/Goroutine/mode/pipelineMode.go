package main

import (
	"fmt"
)

/**
 * 流水线模式 (有先后顺序)
 * 先后顺序：采购 -> 组装 -> 打包   (前一道工序没有完成后一道工序无法进行，存在依赖关系)
 */

// 工序1采购
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

// 工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

// 工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}

func main() {
	// 采购6套配件
	accessories := buy(6)
	// 组装6部电脑
	computers := build(accessories)
	// 打包它们以便售卖
	packs := pack(computers)
	// 输出测试
	for p := range packs {
		fmt.Println(p)
	}
}

// 打包(组装(配件1))
// 打包(组装(配件2))
// 打包(组装(配件3))
// 打包(组装(配件4))
// 打包(组装(配件5))
// 打包(组装(配件6))
