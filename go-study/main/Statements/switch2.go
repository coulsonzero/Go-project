//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	x := 2
	switch {
	case x > 0 && x < 10:
		fmt.Println("ok")
	case x > 10:
		fmt.Println("no")
	}
}
