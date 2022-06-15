package main

import (
	"fmt"
	"runtime"
)

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
	fmt.Println(runtime.Version())
	s := Student{"小明", true, 23, 88, Address{"湖南省", "长沙市"}}
	fmt.Println(s.address.province)

	add := Address{
		city:     "beijing",
		province: "chaoyang",
	}
	str := Student{
		name:    "John",
		sex:     true,
		age:     20,
		address: add,
	}
	fmt.Println(str)

}
