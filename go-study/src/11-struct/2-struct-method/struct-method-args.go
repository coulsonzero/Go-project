package main

import "fmt"

type person struct {
	name string
	age  int
}

// 1. struct 作为 method参数
func hello(p person) {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

// 2. struct 作为 method 接收者
func (p person) change() {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

// pointer -> set field
func (p *person) change_ptr() {
	p.age++
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

// 3. struct 作为 返回类型
func newPerson(name string, age int) *person {
	return &person{name: name, age: age}
}

func main() {
	p := person{"James", 20}

	// 结构体作为函数的参数
	hello(p)
	// `结构体`.方法()
	p.change() // 20
	// 使用指针作为接收者修改结构体的数据
	p.change_ptr() // 21

	p2 := newPerson("tom", 27) // &{tom 27}
	fmt.Println(p2)
}
