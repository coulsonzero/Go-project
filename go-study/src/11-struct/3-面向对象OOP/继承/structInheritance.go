package main

type Animal struct {
	Name string
}

type Cat struct {
	Animal // 继承
}

func (a *Animal) getName(name string) string {
	return a.Name
}

// Method Rewriting (方法名、参数都相同)
func (c *Cat) getName() string {
	// 调用父类 字段/方法
	return c.Animal.Name
}
