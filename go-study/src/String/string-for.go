package main

import "fmt"

func main() {
	c := 'z'
	fmt.Printf("char: %c, value: %v, typeof: %T \n", c, c, c)                         // char: z, value: 122, typeof: int32
	fmt.Printf("char: %c, value: %v, typeof: %T \n", byte(c), byte(c), byte(c))       // char: z, value: 122, typeof: uint8
	fmt.Printf("char: %s, value: %v, typeof: %T \n", string(c), string(c), string(c)) // char: z, value: z,   typeof: string

	z := '美'
	fmt.Printf("char: %c, value: %v, typeof: %T \n", z, z, z)                         // char: 美, value: 32654, typeof: int32
	fmt.Printf("char: %c, value: %v, typeof: %T \n", byte(z), byte(z), byte(z))       // char: *, value: 142, typeof: uint8
	fmt.Printf("char: %s, value: %v, typeof: %T \n", string(z), string(z), string(z)) // char: 美, value: 美, typeof: string
}

func forEach() {
	s := "Github官网"
	for _, v := range s {
		fmt.Printf("%c(%T) ", v, v)
	}
	fmt.Println()
	// G i t h u b 官 网
}

func forEach2() {
	s := "Github官网"
	for _, v := range []rune(s) {
		fmt.Printf("%v(%T) ", string(v), string(v))
	}
	fmt.Println()
	// G i t h u b 官 网
}

func forEach3() {
	s := "Github官网"
	for i := 0; i < len([]rune(s)); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	// 71(G) 105(i) 116(t) 104(h) 117(u) 98(b) 229(å) 174(®)
}
