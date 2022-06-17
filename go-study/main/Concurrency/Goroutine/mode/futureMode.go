package main

import (
	"fmt"
	"time"
)

/**
 * 未来模式：主线程无需等待子线程执行完返回结果，可以先去做其他事情，然后等到未来需要子线程结果的时候再来取
 * 沏茶：放茶叶、烧水，这两个步骤没有依赖关系，是独立的可以同时做。
 */

// 放茶叶
func putInTea() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "茶叶已经放入茶杯～"
	}()
	return vegetables
}

// 烧水
func boilingWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "水已经烧开～"
	}()
	return water
}

func main() {
	teaCh := putInTea()       // 放茶叶
	waterCh := boilingWater() // 烧水
	fmt.Println("已经安排放茶叶和烧水，休息2秒钟～")
	time.Sleep(2 * time.Second)

	fmt.Println("沏茶了，看看茶叶和水好了吗？")
	tea := <-teaCh
	water := <-waterCh
	fmt.Println("准备好了，可以沏茶了：", tea, water)
}

// 已经安排放茶叶和烧水，休息2秒钟～
// 沏茶了，看看茶叶和水好了吗？
// 准备好了，可以沏茶了： 茶叶已经放入茶杯～ 水已经烧开～
