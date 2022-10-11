package main

/**
 * question: GP34 推箱子
 * description: 在二维平面上，有一个人在推箱子，假设从原点 (0, 0) 开始。给出它的移动顺序，判断这个箱子在完成移动后是否在 (0, 0) 处结束。移动顺序由字符串 moves 表示。字符 move[i] 表示其第 i 次移动。推箱子的有效动作有 R（右），L（左），U（上）和 D（下）。
 *              如果推箱子完成所有动作后返回原点，则返回 true。否则，返回 false。
 * input : "UD"
 * output: true
 */
func pushBox(forwards string) bool {
	// write code here
	x, y := 0, 0
	charArr := []byte(forwards)
	for _, v := range charArr {
		switch v {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		default:
			break
		}
	}
	return x == 0 && y == 0
}
