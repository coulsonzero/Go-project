package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) // ["a" "b" "c"]

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))  // moo moo moo

	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToUpper("Gopher"))

	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Hello, Gophers
	fmt.Print(strings.TrimFunc("¡¡¡$6521.123Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) // 6521.123Hello, Gophers

}
