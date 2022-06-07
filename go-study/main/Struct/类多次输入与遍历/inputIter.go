package main

import "fmt"

type Cart struct {
	prices []float32
}

func main() {
	c := Cart{}
	var n int
	fmt.Scanln(&n)

	var num float32
	for i := 0; i < n; i++ {
		fmt.Scanln(&num)
		c.prices = append(c.prices, num)
	}

	c.show()
}

func (x Cart) show() {
	var sum float32 = 0.0
	for _, v := range x.prices {
		sum += v
	}
	fmt.Println(sum)
}
