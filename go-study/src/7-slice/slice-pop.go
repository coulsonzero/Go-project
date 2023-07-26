package main

import "fmt"

func main() {
	slice := []int{1, 3, 2, 5, 6}

	change(slice)
	fmt.Printf("%v\n", slice) // [1 3 2 5 6]

	slice = slice[:len(slice)-1]
	fmt.Printf("%v\n", slice) // [1 3 2 5]
}

func change(slice []int) {
	slice = slice[:len(slice)-1]
}
