package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	s := Student{"Tom", 21}
	s.age = 8
	fmt.Println(s.name) // "Tom"
	fmt.Println(s.age)  // 8
	fmt.Println(s)      // {"Tom", 8}
}
