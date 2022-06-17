package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/**
 * 使用go连接Redis
 * 需先在命令行启动redis服务
 */

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis server:", err)
		return
	}
	fmt.Println("connect redis success!")
	defer conn.Close()
}
