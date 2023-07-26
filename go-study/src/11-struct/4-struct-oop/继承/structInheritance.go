package main

type Animal struct {
	Name string
}

func (a Animal) getName(name string) string {
	return a.Name
}

type Cat struct {
	*Animal // 指针继承
}

// Method Rewriting
func (c Cat) getName() string {
	return c.Animal.Name
}

func (c *Cat) setName(name string) {
	// c.Animal = new(Animal) // 指针继承需要使用new开辟内存空间
	// c.Animal.Name = name
	c.Animal = &Animal{name}
}

func main() {
	// 继承

	// var c Cat
	// c.setName("poul")
	// fmt.Println(c.Name)
	// fmt.Println(c.Animal.Name)
	// fmt.Println(c.getName())

	// 组合
	// c := &Cat{&Animal{Name: "poul"}}
}
