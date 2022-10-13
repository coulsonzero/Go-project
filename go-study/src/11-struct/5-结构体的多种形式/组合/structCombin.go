package main

import "fmt"

type Person struct {
}

type Student struct {
	*Person
}

func (p *Person) GetName(name string) {
	fmt.Printf("Good Morning! %s \n", name)
}

func main() {
	p := Person{}
	p.GetName("John") // Good Morning! John

	// 组合
	s := Student{&Person{}}
	s.GetName("John") // Good Morning! John
}
