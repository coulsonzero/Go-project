package main

import "fmt"

func main() {

	c := 'a'
	z := '美'

	fmt.Printf("char: %c, value: %v, typeof: %T \n", c, c, c)                         // char: a, value: 97, typeof: int32
	fmt.Printf("char: %c, value: %v, typeof: %T \n", byte(c), byte(c), byte(c))       // char: a, value: 97, typeof: uint8
	fmt.Printf("char: %s, value: %v, typeof: %T \n", string(c), string(c), string(c)) // char: a, value: a,   typeof: string

	fmt.Printf("char: %c, value: %v, typeof: %T \n", z, z, z)                         // char: 美, value: 32654, typeof: int32
	fmt.Printf("char: %c, value: %v, typeof: %T \n", byte(z), byte(z), byte(z))       // char: *, value: 142, typeof: uint8
	fmt.Printf("char: %s, value: %v, typeof: %T \n", string(z), string(z), string(z)) // char: 美, value: 美, typeof: string

}
