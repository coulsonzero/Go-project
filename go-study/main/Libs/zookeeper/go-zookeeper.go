package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

/**
 * Zookeeper: 分布式锁
 * 优点：具备高可用、可重入、阻塞锁的特性，可解决失效死锁的问题
 * 缺点：由于需要频繁地创建和删除节点，性能上不如Redis方式
 * Zookeeper客户端：Curator
 */

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) // *10)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
