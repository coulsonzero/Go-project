package main

import "fmt"

//go:generate go run main.go
//go:generate go version
//go:generate ls
func main() {
	fmt.Println("hello world!")
}

// $ go generate -x
// go run main.go
// hello world!
// go version
// go version go1.19.3 darwin/arm64
// ls
// main.go
