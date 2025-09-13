package main


import (
"fmt"
)

func main() {
	ret := test2()
	fmt.Println("test return:", ret)
}

//返回值改为命名返回值
func test() (i int) {
	//var i int

	defer func() {
		i++
		fmt.Println("test defer, i = ", i)
	}()

	return i
}

func test2() (int) {
	var i int

	defer func() {
		i++
		fmt.Println("test defer, i = ", i)
	}()

	return i
}