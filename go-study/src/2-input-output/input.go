package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Print("please enter your name: ")
	fmt.Scanln(&name)
	// fmt.Println("Output: " + name)

	// bufio_example()
}

// please enter your name: John Smith
// Output: John

func bufio_example() {
	ans, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	print(ans)
}

// John Smith
// John Smith



