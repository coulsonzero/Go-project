package main

import "fmt"

func main() {
	fmt.Println(sumInt(2, 3))
	fmt.Println(sumFloat(1.2, 1.3))
	fmt.Println(sum[float64](1.2, 1.3))
	fmt.Println(sum[int](1, 3))
	fmt.Println(Sum(1, 3))
	fmt.Println(Sum(1.2, 3.4))

}

func sumInt(a, b int) int {
	return a + b
}

func sumFloat(a, b float64) float64 {
	return a + b
}

func sum[T int | float64](a, b T) T {
	return a + b
}

type Num interface {
	int | float64
}

func Sum[T Num](a, b T) T {
	return a + b
}
