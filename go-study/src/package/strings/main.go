package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func main() {
	fmt.Println(strings.HasSuffix("main.py", "py"))
	fmt.Println(strings.HasPrefix("main.py", "py"))
}

func join() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}

func split() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) // ["a" "b" "c"]
}

// 包含
func contains() {
	if ok := strings.Contains("test-v1", "v1"); ok {
		fmt.Println("find the character.")
		fmt.Println(ok)
	}
}

func replace() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))  // moo moo moo
}

func demo() {
	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.ToUpper("Gopher"))

	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Hello, Gophers
	fmt.Println(strings.TrimFunc("¡¡¡$6521.123Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) // 6521.123Hello, Gophers
}

func progressBar() {
	const col = 30
	// Clear the screen by printing \x0c.
	bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
}
