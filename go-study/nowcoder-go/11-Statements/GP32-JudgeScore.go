package main

/**
 * question: GP32 成绩判定
 * description: 根据成绩分数输出成绩等级，判定规则如下 分数低于60 算不及格，60-80 含60 为中等， 80-90含80  为良好，90分以上含90 为优秀。
 * input : 59
 * output: "不及格"
 */
func judgeScore(score int) string {
	// write code here
	switch {
	case score < 60:
		return "不及格"
	case score >= 60 && score < 80:
		return "中等"
	case score >= 80 && score < 90:
		return "良好"
	case score >= 90:
		return "优秀"
	default:
		return ""
	}
}
