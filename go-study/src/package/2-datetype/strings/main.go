package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	// byteApi()
	// typeConvert()
	// judgeApi()

	// println(strings.Count("abcaabc12", "a"))
}

func byteApi() {
	// 字母
	println(unicode.IsLetter('a')) // true
	println(unicode.IsUpper('A'))  // true
	println(unicode.IsLower('a'))  // true
	// 数字, 判断的严格性：c >= '0' && c <= '9' ⇒ IsDigit ⇒ IsNumber
	c := byte('7')
	println(c >= '0' && c <= '9')  // true,	['1-9']
	println(unicode.IsDigit('7'))  // true, ['1-9']
	println(unicode.IsNumber('½')) // true, ['1-9', 'Ⅷ', '½']
	// 空白字符
	println(unicode.IsSpace('\n')) // true, [' ', '\n', '\t']
	// 标点字符
	println(unicode.IsPunct(',')) // true
}

func typeConvert() {
	println(strconv.Atoi("12"))             // string -> number(int)
	println(strconv.ParseInt("12", 10, 64)) // string -> number(int64)
	println(strconv.Itoa(12))               // number -> string
	println(fmt.Sprintf("%d", 12))          // number -> string

	// string <-> array
	println(strings.Join([]string{"127", "0", "0", "1"}, "."))
	fmt.Printf("%v \n", strings.Split("127.0.0.1", "."))
}

func stringApi() {
	// +
	println("hello" + ", " + "world")

	// build | buffer
	var buf strings.Builder
	// var buf bytes.Buffer
	buf.WriteString("hello")
	buf.WriteByte(',')
	buf.Write([]byte("world"))
	println(buf.String())
}

func judgeApi() {

	println(strings.Contains("test-v1", "v1"))
	// startsWith, endsWith
	println(strings.HasPrefix("main.go", "main"))
	println(strings.HasSuffix("main.go", "go"))

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
