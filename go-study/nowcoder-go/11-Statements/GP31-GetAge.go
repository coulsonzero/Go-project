package main

/**
 * question: GP31 年龄分段
 * description: 已知年龄的分段如下， 婴儿(出生0-1岁)、幼儿(1-4岁)包含1岁、儿童(5-11)包含5岁、少年(12-18)包含12岁、青年(19-35)包含19岁、中年(36-59)包含36岁、老年(60以上)包含60岁 ，
 *              输入一个人的年龄，返回相应的年龄段。
 * input : 35
 * output: "青年"
 */
func getAge(age int) string {
	// write code here
	switch {
	case age >= 0 && age < 1:
		return "婴儿"
	case age >= 1 && age < 5:
		return "幼儿"
	case age >= 5 && age < 12:
		return "儿童"
	case age >= 12 && age < 19:
		return "少年"
	case age >= 19 && age < 36:
		return "青年"
	case age >= 36 && age < 60:
		return "中年"
	case age >= 60:
		return "老年"
	default:
		return ""
	}
}
