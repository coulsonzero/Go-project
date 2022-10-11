package main

import "fmt"

type Student struct {
	name string
	age  int
}

func hello(s Student) {
	fmt.Printf("name: %s, age: %d\n", s.name, s.age)
}

func (s Student) welcome() {
	fmt.Printf("name: %s, age: %d\n", s.name, s.age)
}

func (s *Student) change_ptr() {
	s.age += 1
	fmt.Printf("name: %s, age: %d\n", s.name, s.age)
}

func main() {
	s := Student{"James", 20}

	// 结构体作为函数的参数
	hello(s)
	// `结构体`.方法()
	s.welcome()
	// 使用指针作为接收者修改结构体的数据
	s.change_ptr()
}
