package main

import (
	"fmt"
)

/**
 * question: GP44 动物和老虎
 * description: 定义一个动物接口，该接口有 sleep,eat 方法，定义老虎实现动物接口，老虎的sleep方法输出"tiger sleep",eat方法输出"tiger eat"，最后依次调用老虎的sleep,eat方法。
 * input :
 * output: sleep
 *         eat
 */

type Animal interface {
	eat()
	sleep()
}
type Tiger struct{}

func (t Tiger) eat() {
	fmt.Println("eat")
}

func (t Tiger) sleep() {
	fmt.Println("sleep")
}

func main() {
	var t Animal = Tiger{}
	t.sleep()
	t.eat()
}
