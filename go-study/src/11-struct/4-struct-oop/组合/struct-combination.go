package main

import "fmt"

type Person struct{}

func (p *Person) GetName(name string) {
	fmt.Printf("Good Morning! %s \n", name)
}

type Student struct {
	*Person
}

func main() {
	// 组合：继承基础上, 实例化子结构体时内部初始化父结构体，类似包含关系
	s := &Student{&Person{}}
	s.GetName("John") // Good Morning! John
}
