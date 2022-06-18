package main

import "fmt"

func main() {
	if x := 42; x > 18 {
		fmt.Println("access")
	} else if x == 18 {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}

	// fmt.Println(x) // error: undefined: x
}
