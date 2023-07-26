package main

func main() {
	s := "Hello World !"

	// fmt.Printf("%v \n", &s)
	//
	// s = "title: lop"
	//
	// fmt.Printf("%v \n", &s)
	// fmt.Println(s)

	[]byte(s)[0] = 'a'

}
