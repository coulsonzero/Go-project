package main

/**
 * question: GP37 质量检查
 * description: 某工厂加工了一批材料，现将加工出来的成品抽出来一部分进行检查，给定一个标准，如果低于这个标准，则视为不合格。所有成品的质量用一个切片来表示。
 * input : [2,2,3,4,6,6,3],3
 * output: [3,4,6,6,3]
 */
func check(material []int, standard int) []int {
	// write code here
	res := []int{}
	for _, v := range material {
		if v >= standard {
			res = append(res, v)
		}
	}
	return res
}
