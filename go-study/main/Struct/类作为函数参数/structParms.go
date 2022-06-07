package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s Student) welcome() {
	fmt.Println(s.name)
	fmt.Println(s.age)
}

func (p Student) set(val int) {
	p.age += val
}

func (ptr *Student) set_ptr(val int) {
	ptr.age += val
}

func main() {
	// 方式一
	s := Student{"Jame", 20}
	s.welcome() // "Jame", 20

	s.set(3)
	fmt.Println(s.age) // 20
	s.set_ptr(4)
	fmt.Println(s.age) // 24

	// 方式二：指针
	s2 := Student{"Tom", 16}
	ptr2 := &s2

	fmt.Println(*ptr2) // {"Tom", 16}
	ptr2.name = "Coulson"
	fmt.Println(s2) // {"Coulson", 16}
	ptr2.set(4)
	fmt.Println(s2) // {"Coulson", 16}
	ptr2.set_ptr(4)
	fmt.Println(s2) // {"Coulson", 20}

	// 方式三：内存地址
	s3 := &Student{"Yalo", 16}
	fmt.Println(*s3)
	s3.set(4)
	fmt.Println(*s3) //{Yalo 16}
	s3.set_ptr(4)
	fmt.Println(*s3) // {Yalo 20}
}
