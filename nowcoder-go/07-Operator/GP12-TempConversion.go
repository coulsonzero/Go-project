package main

/**
 * question: GP12 温度转换
 * description: 定义一个变量保存华氏温度，华氏温度转换摄氏度的公式为：5/9*(华氏温度-32），求出华氏温度对应的摄氏温度。
 * input: 100.0
 * output: 37.77778
 */

func temperature(f float64) float64 {
	// write code here
	return 5 * (f - 32.0) / 9
}
