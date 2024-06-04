package main

import "fmt"


func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}


func typeof(t any) {
	switch t := t.(type) {
	case bool:
		fmt.Printf("boolean %t\n", t) 		  // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) 		  // t has type int
	case float64, float32:
		fmt.Printf("float   %f\n", t)
	case string:
		fmt.Printf("string  %s\n", t)
	case byte, rune:
		fmt.Printf("char    %c\n", t)
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	}
}


func main() {
	fmt.Println(shouldEscape('-')) 		// false
	fmt.Println(shouldEscape('#')) 		// true
	fmt.Println(shouldEscape('+')) 		// true

	typeof(12)     		 // integer
	typeof(12.3)         // float
	typeof(false)   	 // boolean
	typeof("hello") 	 // string
	typeof('k') 	 	 // char

}


