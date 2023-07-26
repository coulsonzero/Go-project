package encapsulation

import "fmt"

func main() {
	p := Person{}
	p.SetName("John")
	fmt.Println(p)

	s := Student{}
	fmt.Println(s.getGender())
}
