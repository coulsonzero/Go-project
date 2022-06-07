package main

import "fmt"

type Student struct {
	name string
	age  int
}

func resetName(ptr *Student) {
	ptr.name = "Coulson"
}

func resetAge(p *Student) {
	p.age = 0
}

func main() {
	s := Student{"Tom", 21}
	p := &s

	fmt.Println(p.name) // "Tom"
	fmt.Println(p.age)  // 21
	fmt.Println(*p)     // {"Tom", 21}

	resetAge(&s) // {"Tom", 0}
	fmt.Println(s)

	resetName(p)
	fmt.Println(*p) //{"Coulson", 0}

	s2 := &Student{"Jome", 23}
	fmt.Println(*s2)     // {"Jame", 23}
	fmt.Println(s2.name) // "Jame"
	fmt.Println(s2.age)  // 23
}
