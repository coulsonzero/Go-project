package main

import "fmt"

func main() {
	x := 8
	switch y := x % 2; y {
	case 1:
		fmt.Println("1")
		//statement(s)        //不需要break
	default:
		fmt.Println("0")
		//statement(s)
	}
}
