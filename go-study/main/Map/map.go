package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jame"] = 42
	m["Amye"] = 24

	fmt.Println(m) // map[Amye:24 Jame:42]

}
