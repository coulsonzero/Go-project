package main

/**
 * question: GP21 出队
 * description: 学生们都排成了一队，有一个切片表示相应学生们的身高，现随机喊某个位置的人出队，返回出队后的这个切片。比如[2,3,4,5],索引为1的位置的出队，出队后切片为[2,4,5]
 * input : [1,2,3,4,5,6],3
 * output: [1,2,3,5,6]

 */
func deleteElement(s []int, index int) []int {
	// write code here
	return append(s[:index], s[index+1:]...)
}
