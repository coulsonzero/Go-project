package main

/**
 * question: GP36 坐标转换
 * description: 已知一个m*n二维数组，二维数组中的元素的索引（x，y）可以表示为一个二维坐标,现将这个二维坐标转换为一维坐标,一维坐标=x*n+y。返回这个一维数组。
 * input : [[1,2,3],[4,5,6],[7,8,9]]
 * output: [1,2,3,4,5,6,7,8,9]
 */
func transform(array [][]int) []int {
	// write code here
	res := []int{}
	for _, v := range array {
		res = append(res, v...)
	}
	return res
}
