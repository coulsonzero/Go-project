package main

import "fmt"

type Cat struct {
	name string
	age  int
}

func main() {
	cat := Cat{name: "pat"}

	// cat := new(Cat)
	// var cat Cat
	// cat.name = "pat"

	fmt.Println(cat)
}






