package main

import "fmt"

func main() {
	c := swapCase('a')
	fmt.Printf("%c \n", c) // A

}

func swapCase(c byte) byte {
	c ^= 32
	return c
}
