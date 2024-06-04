package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	println(strconv.Atoi("12")) // 12
	println(strconv.Itoa(32))   // "32"

	println(unicode.IsLetter('f')) // true
	println(unicode.IsDigit(7))    // false
	println(unicode.IsDigit('7'))  // true
	println(unicode.IsSpace('\n')) // true

	fmt.Println(string('H' + 32)) // 'H' -> 'h'
	fmt.Println(string('p' - 32)) // 'p' -> 'P'
}

func stringIter() {
	s := "hello, world!"
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
	}

	for _, v := range s {
		print(string(v))
	}
}

func stringBuild() {
	var buf strings.Builder // buf := new(strings.Builder)
	// var buf bytes.Buffer

	buf.WriteString("hello")
	buf.WriteByte(',')
	buf.WriteString("world")
	buf.WriteRune('!')
	println(buf.String()) // hello,world!
}

func stringApi() {
	s := "hello world"
	println(len(s)) // 11
	for i := 0; i < len(s); i++ {
		print(string(s[i]))
	}

	println(strings.Index(s, "wo"))          // 6
	println(strings.Index(s, "w"))           // 6
	println(strings.ReplaceAll(s, " ", ",")) // hello,world

	fmt.Println(strings.Split("hello, world!", ", ")) // ["hello", "world!"]
	slice := []string{"hello", "world"}
	fmt.Println(strings.Join(slice, ",")) // "hello, world!"

	println(strings.Contains("hello.go", "l"))   // true
	println(strings.HasPrefix("hello.go", "he")) // true
	println(strings.HasSuffix("hello.go", "go")) // true

	println(strings.ToUpper("Hello"))                          // HELLO
	println(strings.ToLower("Hello"))                          // hello
	println(strings.Title("hello"))                            // Hello
	println(cases.Title(language.Und).String("hello, world!")) // Hello, World!
}

func stringInit() {
	// var s string
	// s = "Hello World"

	// s := "hello world"

	// var s string = "hello world!"

	name := "coulson"
	s := fmt.Sprintf("hello %s", name)
	fmt.Println(s)
}
