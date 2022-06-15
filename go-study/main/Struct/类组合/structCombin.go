package main

import "fmt"

type Person struct {
}

func (p *Person) GetName(name string) {
	fmt.Printf("Good Morning! %s \n", name)
}

type Student struct {
	*Person
}

func main() {
	p := Person{}
	p.GetName("John") // Good Morning! John

	s := Student{&Person{}}
	s.GetName("John") // Good Morning! John
}
