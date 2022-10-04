package main

import (
	"fmt"
)

// type any = interface{}

func main() {
	Hello("coulson")
}

func Hello(name string) {
	print := func(a ...interface{}) {
		fmt.Printf("Hello %s", a...)
	}

	// if name == "" {
	// 	print("Name cannot be empty！")
	// 	return
	// }
	print(name)
}
