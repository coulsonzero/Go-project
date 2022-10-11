// new & make
// new: 分配一个类型的内存
// make: 仅用于对象类型(slice, map, chan)的内存分配
// new: 返回新内存地址的指针
// make: 返回第一个参数类型的引用，而不是指针

package main

import "fmt"

func newDemo() {
	var i *int = new(int)
	fmt.Println(i) // 0x14000122008

	var slice = new([]int)
	fmt.Println(slice) // &[]
}

func makeDemo() {
	var slice = make([]int, 3, 5)
	fmt.Println(slice) // [0 0 0]

	var m = make(map[int]string)
	fmt.Println(m) // map[]

	var ch = make(chan int)
	fmt.Println(ch) // 0x14000102060

	var ch_bf = make(chan int, 2)
	fmt.Println(ch_bf) // 0x140000ba000
}
