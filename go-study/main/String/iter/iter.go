package main

import "fmt"

func main() {
	s := "hello world"

	for _, c := range s {
		fmt.Println(c)
	}
}
