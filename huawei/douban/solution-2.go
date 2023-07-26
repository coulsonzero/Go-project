package main

func main() {
	println(isPatten("x+1yz*z", "xy1yzxyyz"))  // true
	println(isPatten("x+1yz+z", "xy1yzxyyz"))  // false
	println(isPatten("x+1yz++z", "xy1yzxyyz")) // false
	println(isPatten("x+1yz+*z", "xy1yzxyyz")) // true
	// println(isPatten("x+1yz*+z", "xy1yzxyyz")) // error

}

func isPatten(p, s string) bool {
	i := 0
	j := 0
	for i < len(p) && j < len(s) {
		if p[i] == s[j] || p[i] == '+' {
			i++
			j++
		} else if p[i] == '*' {
			if s[j+1] == p[i+1] {
				i++
				j++
			} else {
				j++
			}
		} else {
			break
		}
	}

	// fmt.Println(i, j)
	// fmt.Println(len(p), len(s))
	return len(p) == i && len(s) == j
}
