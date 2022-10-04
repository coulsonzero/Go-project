package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	s := &student{name: "john", age: 20}

	fmt.Println(*s)        // {john 20}
	fmt.Println(s)         // &{john 20}
	fmt.Println(s.name)    // john
	fmt.Println((*s).name) // john
}
