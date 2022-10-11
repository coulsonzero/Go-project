// fmt
// fmt.Print()
// fmt.Printf()
// fmt.Println()
// fmt.Sprintf()

package main

import (
	"fmt"
	"os"
)

func main() {
	printColor()

}

func print() {
	fmt.Print("hello\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
}

func printStruct() {
	type T struct {
		a int
		b float64
		c string
	}
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
}

func printColor() {
	fmt.Printf("\n %c[0;34;34m%s%c[0m\n\n", 0x1B, "testPrintColor", 0x1B)
}
