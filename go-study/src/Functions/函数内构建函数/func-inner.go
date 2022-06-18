package main

import (
	"fmt"
)

func main() {
	Hello("coulson")
}

func Hello(name string) {
	print := func(a ...any) {
		fmt.Printf("Hello %s", a...)
	}

	if name == "" {
		print("Name cannot be empty！")
		return
	}
	print(name)
}
