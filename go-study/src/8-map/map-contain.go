package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["tom"] = 20
	m["john"] = 18

	fmt.Println(m["uo"]) // 0

	if age, ok := m["tom"]; ok {
		fmt.Println(age) // 20
	}
	if age, ok := m["uo"]; !ok {
		fmt.Println(age) // 0
	}
}
