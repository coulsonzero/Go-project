package main

import "fmt"

func main() {
	ArrayInit()

}

func ArrayInit() {
	var a [3]int                             // len: 3, cap: 3, array: [0 0 0], array: [3]int{0, 0, 0}
	var b = [...]int{1, 2, 3}                // len: 3, cap: 3, array: [1 2 3]
	var c = [...]int{1: 2, 2: 3}             // len: 3, cap: 3, array: [0 2 3]
	var d = [...]int{1, 2, 4: 5, 6}          // len: 6, cap: 6, array: [1 2 0 0 5 6]
	var e = [...]int{'a': 1}                 // len: 98, cap: 98, array: [..., 1]
	var f = [...]int{'a': 1, 'b': 2, 'c': 3} // len: 100, cap: 100, array: [..., 1 2 3]

	fmt.Printf("len: %d, cap: %d, array: %v, array: %#v \n", len(a), cap(a), a, a)
	fmt.Printf("len: %d, cap: %d, array: %v \n", len(b), cap(b), b)
	fmt.Printf("len: %d, cap: %d, array: %v \n", len(c), cap(c), c)
	fmt.Printf("len: %d, cap: %d, array: %v \n", len(d), cap(d), d)
	fmt.Printf("len: %d, cap: %d, array: %v \n", len(e), cap(e), e)
	fmt.Printf("len: %d, cap: %d, array: %v \n", len(f), cap(f), f)

	fmt.Printf("%v, %c, %T", 'a', 'a', 'a') // 97, a, int32

	demo()
}

func ArrayIter() {
	var nums = [5]int{1, 2, 3, 4, 5}
	ptr := &nums
	for _, v := range ptr {
		fmt.Println(v)
	}
}

func demo() {
	var p [100]int
	var m interface{} = [...]int{99: 0}
	// fmt.Println(p)
	// fmt.Println(m)
	fmt.Println(p == m) // true
}
