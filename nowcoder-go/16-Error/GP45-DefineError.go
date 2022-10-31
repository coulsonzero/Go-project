package main

/**
 * question: GP45 网络延迟
 * description: 实现erro接口，自定义一个错误，该错误抛出"网络延迟"错误。输入网络的延迟数，如果延迟数大于100则认为网络延迟，并返回
 * input : 150
 * output: "网络延迟"
 */

type error interface {
	Error() string
}

type errors struct {
}

func (e errors) Error() string {
	return "网络延迟"
}

func defineerr(ping int) string {
	// write code here
	if ping > 100 {
		// return errors.New("网络延迟").Error()
		return errors{}.Error()
	}
	return ""
}

func main() {
	println(defineerr(200))
}
