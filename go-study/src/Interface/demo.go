package main

import (
	"fmt"
)

type Animal interface {
	sleep()
	eat()
}
type Tiger struct{}

func (t Tiger) sleep() {
	fmt.Println("sleep")
}

func (t Tiger) eat() {
	fmt.Println("eat")
}

func main() {
	var t Animal = Tiger{}
	t.sleep()
	t.eat()
}
