package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender bool
	Score  float64
}

func main() {
	// 方式一
	var person Person
	person.Name = "John"
	person.Age = 20
	fmt.Println(person)
	// Output: {John 20 false 0}

	// 方式二
	var p3 = new(Person)
	p3.Name = "John"
	(*p3).Age = 20
	fmt.Println(*p3)
	// Output: {John 20 false 0}

	// 方式三：必须要写全！
	p1 := Person{"John", 20, true, 97.2}
	fmt.Println(p1)
	// Output: {John 20 true 97.2}

	// 方式四(推荐)
	p2 := Person{Name: "John", Age: 20}
	p2.Score = 97.2
	fmt.Println(p2)
	// Output: {John 20 false 97.2}

	// 方式五(推荐)
	p4 := &Person{Score: 97.2}
	// var p4 = &Person{Score: 97.2}
	p4.Name = "John"
	(*p4).Age = 20
	fmt.Println(*p4)
	// Output: {John 20 false 97.2}
}
