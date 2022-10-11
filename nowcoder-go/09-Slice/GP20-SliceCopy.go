package main

/**
 * question: GP20 切片复制
 * description: 给定一个切片和另一个空切片，将第一个切片复制到第二个空切片中，并返回这个被复制的空切片
 * input : [1,2],[]
 * output: [1,2]
 */
func sliceCopy(src []int, des []int) []int {
	// write code here
	des = src[:]
	return des
}
