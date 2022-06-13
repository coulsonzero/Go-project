package main

import "fmt"

func main() {
	// s := "hello world"

	// for _, c := range s {
	// 	fmt.Println(c)
	// }
	// traversalString()

}

func traversalString() {
	s := "Github官网"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	// 71(G) 105(i) 116(t) 104(h) 117(u) 98(b) 229(å) 174(®) 152(nul) 231(ç) 189(½) 145(nul)

	for _, r := range s { // rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	// 71(G) 105(i) 116(t) 104(h) 117(u) 98(b) 23448(官) 32593(网)
}
