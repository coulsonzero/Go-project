package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

// 无法修改示例结构体的值
func (s student) change(val int) {
	s.age = val
}

// 通过指针修改结构体的值
func (s *student) change_ptr(val int) {
	s.age = val
}

func main() {
	s := &student{name: "john"}

	fmt.Println(*s)        // {john 20}
	fmt.Println(s)         // &{john 20}
	fmt.Println(s.name)    // john
	fmt.Println((*s).name) // john

	s.change(10)
	fmt.Println(s.age) // 20， 不变
	s.change_ptr(12)
	fmt.Println(s.age) // 12
}
