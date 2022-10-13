package 封装

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) ToString() string {
	return fmt.Sprintf("name: %s, age: %d \n", p.Name, p.Age)
}
