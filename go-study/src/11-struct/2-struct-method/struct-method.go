package main

import "fmt"

type Student struct {
	name string
	age  int
}

// 1. struct as receiver argument
func (s *Student) setName(name string) {
	s.name = name
}

// 2. struct as function argument
func setName(s *Student, name string) {
	s.name = name
}

// 3. struct as return argument
func newStudent(name string) *Student {
	return &Student{name: name}
}

func main() {
	// 1. receiver argument example:
	var s Student
	s.setName("poul")
	fmt.Println(s) // {poul 0}

	// 2. function argument example:
	var s2 Student
	setName(&s2, "koup")
	fmt.Println(s2) // {koup 0}

	// 3. return argument example:
	s3 := newStudent("yerc")
	fmt.Println(*s3) // {yerc 0}
}















