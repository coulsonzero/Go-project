package main

import "fmt"

func main() {
	m := [...]int{1, 2, 3}
	fmt.Println(m)
	constArray()
}

func constArray() {
	m := [...]int{'a': 1, 'b': 2, 'c': 3}
	fmt.Println(m)
	m['a'] = 3
	fmt.Printf("type: %T, len: %d, value: %v", m, len(m), m)
}







