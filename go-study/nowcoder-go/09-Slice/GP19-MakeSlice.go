package main

/**
 * question: GP19 创建切片
 * description: 创建一个制定长度，容量的int类型切片，设置该切片的每个位置的值等于其索引值,最后返回该切片
 * input : 5,5
 * output: [0,1,2,3,4]
 */
func makeslice(length int, capacity int) []int {
	// write code here
	slice := make([]int, length, capacity)
	for i, _ := range slice {
		slice[i] = i
	}
	return slice
}
