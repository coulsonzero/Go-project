package main

import "fmt"

func main() {
	cat := Cat{}
	cat.getName("Tim")

}

type Animal struct {
	Name string
}

// private method
func (a Animal) getName(name string) {
	fmt.Printf("Good Morning! %s \n", name)
}

type Cat struct {
	Animal // 继承
}

// Method Rewriting
// 子类重写父类方法(方法名、参数都相同)
func (c Cat) getName(name string) {
	// 调用父类 字段/方法
	fmt.Println(c.Animal.Name)
	c.Animal.getName(name)

	fmt.Printf("Nice to meet you! %s \n", name)
}
