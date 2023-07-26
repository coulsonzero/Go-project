package 递归

func jumpFloor(number int) int {
	// write code here
	a, b := 1, 1
	for i := 1; i < number; i++ {
		a, b = b, a+b
	}
	return b
}
