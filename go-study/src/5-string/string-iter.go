package main

import "fmt"

/*
for i := 0; i < len(s); i++ {
	// s[v] : byte
}

for _, v := range s {
	// v: rune
}
*/

func main() {
	s1 := "hello world"
	s2 := "Github官网"

	// traversalString1(s1) // hello world
	// traversalString2(s1) // hello world
	traversalString3(s1) // hello world
	// traversalString4(s1) // hello world

	// traversalString1(s2) // Githubå®®ç½
	// traversalString2(s2) // Githubå®
	traversalString3(s2) // Github官网
	// traversalString4(s2) // Github官网
}

// func traversalString1(s string) {
// 	for i := 0; i < len(s); i++ {
// 		fmt.Print(string(s[i]))
// 	}
// 	fmt.Println()
// }

// func traversalString2(s string) {
// 	for i := 0; i < len([]rune(s)); i++ {
// 		fmt.Print(string(s[i]))
// 	}
// 	fmt.Println()
// }

func traversalString3(s string) {
	for _, v := range s { // rune
		fmt.Printf(string(v))
	}
	fmt.Println()

}

// func traversalString4(s string) {
// 	for _, v := range []rune(s) {
// 		fmt.Printf(string(v))
// 	}
// 	fmt.Println()
// }
