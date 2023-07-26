package main

import "strconv"

func addBinary(a string, b string) string {
	println(BinaryToInt(a))
	println(BinaryToInt(b))
	println(BinaryToInt(a) + BinaryToInt(b))

	ans := BinaryToInt(a) + BinaryToInt(b)

	return strconv.FormatInt(int64(ans), 2)
}

func BinaryToInt(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

func main() {
	// a := "1010"
	// b := "10111"
	// println(addBinary(a, b))
	println(BinaryToInt("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"))
	println(BinaryToInt("110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))

	println(BinaryToInt("110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"))
}
