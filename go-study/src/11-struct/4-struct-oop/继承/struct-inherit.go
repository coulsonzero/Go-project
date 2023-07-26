package main

import "fmt"

/**
 * 继承的意义：
 * 子结构体能够继承父结构体的 methods 和 fields, 既能重写该方法，也能添加父结构体中没有的新方法
 * 1. 通过 'this.父结构体.字段名' 调用父结构体字段
 * 2. 通过 'this.父结构体.方法名' 调用父结构体方法
 * 3. 通过'this.父结构体 = &父结构体{}' 实例化父结构体
 */

type Father struct {
	Name string
}

func (this *Father) Work() {
	fmt.Println("I'm working...")
}

type Child struct {
	*Father // inherit
}

// method rewriting
func (c *Child) Work() {
	// use the method of the father struct

	// this.Father.Work()
	// fmt.Println("child is working...")

	c.Father = &Father{"poil"}
}

func main() {
	child := Child{}
	child.Work()
	fmt.Println(child.Name)
}
