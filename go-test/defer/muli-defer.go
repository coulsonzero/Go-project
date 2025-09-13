package main

import "fmt"

func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	defer fmt.Println("d")
	fmt.Println("d")
	defer e()
	fmt.Println("f")
}

func e() {
	fmt.Println("e")
}


/*
a
d
f
e
d
c
b
 */

