package main

import "fmt"

type Address struct {
	city     string
	province string
}

type Student struct {
	name    string
	sex     bool
	age     int
	score   int
	address Address
}

func main() {
	s := Student{"小明", true, 23, 88, Address{"湖南省", "长沙市"}}
	fmt.Println(s.name)
	fmt.Println(s.sex)
	fmt.Println(s.age)
	fmt.Println(s.score)
	fmt.Println(s.address.city)
	fmt.Println(s.address.province)
}
