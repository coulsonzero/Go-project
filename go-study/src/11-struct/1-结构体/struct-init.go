package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// 方式一
	var p1 person
	p1.name = "John"
	p1.age = 20

	// 方式二
	p2 := new(person)
	p2.name = "John"
	(*p2).age = 20

	// 方式三：必须要写全！
	p3 := person{"John", 20}
	fmt.Println(p3)

	// 方式四(推荐)
	p4 := person{name: "John", age: 20}
	fmt.Println(p4)

	// 方式五(推荐)
	p5 := &person{name: "John"}
	(*p5).age = 20

	// 方式六
	p6 := newPerson("John", 20)
	fmt.Println(p6)
}

// 需要使用指针
func newPerson(name string, age int) *person {
	return &person{name: name, age: age}
}
