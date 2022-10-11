package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := &Person{"John", 20}
	// 匿名结构体
	student := struct {
		score   float64
		class   int
		persons Person
	}{
		97.2,
		7,
		*person,
	}
	fmt.Printf("%+v", student)
	// Output: {score:97.2 class:7 persons:{Name:John Age:20}}
}
