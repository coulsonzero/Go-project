package main

import "fmt"

type animal interface {
	sleep()
	eat()
}

type cat struct {
}

type dog struct {
}

func (c *cat) sleep() {
	fmt.Println("cat is sleeping ...")
}

func (c *cat) eat() {
	fmt.Println("cat is eating ...")
}

func (d *dog) sleep() {
	fmt.Println("dog is sleeping ...")
}

func (d *dog) eat() {
	fmt.Println("dog is eating ...")
}

func main() {
	var c animal = &cat{}
	c.sleep()
	c.eat()

	var d animal = &dog{}
	d.sleep()
	d.eat()
}

/*
cat is sleeping ...
cat is eating ...
dog is sleeping ...
dog is eating ...
*/
