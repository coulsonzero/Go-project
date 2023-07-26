package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.GC：手动触发 GC
// runtime.ReadMemStats()：读取内存相关的统计信息，其中包含部分 GC 相关的统计信息
// runtime.NumGoroutine(): 查看当前goroutine的数量
// debug.FreeOSMemory：手动将内存归还给操作系统
// debug.ReadGCStats：读取关于 GC 的相关统计信息
// debug.SetGCPercent：设置 GOGC 调步变量
// debug.SetMaxHeap（尚未发布[10]）：设置 Go 程序堆的上限值

func main() {
	// fmt.Println(runtime.Version()) // go1.17.7
	gc()
}

func gc() {
	go func() {
		fmt.Println(runtime.Version())
	}()
	time.Sleep(time.Millisecond)
	runtime.GC()
	println("OK")

}
