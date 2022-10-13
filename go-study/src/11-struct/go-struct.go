package main

import "fmt"

// Struct 结构体
// field: 实例字段的修改需要在方法中的结构体接收者使用指针类型
// method: 结构体可作为参数或接收者
// private: 首字母小写
// public: 首字母大写--> struct名改大写，字段名大写, method名大写

type person struct {
	name string
	age  int
}

// 需要使用指针
func (p *person) setName(name string) {
	p.name = name
}

// 可以不用指针
func (p person) getName() string {
	return p.name
}

// 需要使用指针
func newPerson(name string, age int) *person {
	return &person{name: name, age: age}
}

// 需要使用指针
func getAge(p *person) int {
	return p.age
}

func main() {
	// p := person{name: "shville", age: 21}
	// p.setName("tom")
	// println(p.getName())

	p := newPerson("john", 23)
	p.setName("tom")
	println(p.getName())
	println(getAge(p))
	p.name = "kio"
	fmt.Println(p)
}
