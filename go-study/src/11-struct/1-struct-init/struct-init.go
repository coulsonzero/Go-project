package main

type person struct {
	name string
	age  int
}

func main() {
	// 方式一: value init
	var p person
	p.name = "John"
	p.age = 20

	p3 := person{"John", 20}            // 方式三：必须要写全！
	p4 := person{name: "John", age: 20} // 方式四：字段初始化，相对比较灵活

	// 方式二: pointer init
	p2 := new(person)
	(*p2).name = "John"
	(*p2).age = 20 // '(*p).age' allows to shorten use 'p.age' instead.

	p5 := &person{name: "John", age: 20} // 方式五：返回的是指针
}



