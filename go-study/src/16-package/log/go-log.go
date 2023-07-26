package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("%s", time.Now())
	// 设置日志输出格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 设置日志前缀
	log.SetPrefix("[gin] ")
	log.Println()
	// log.Fatal(err)

}
