package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	length float64
}

type Triangle struct {
	bottom float64
	height float64
}

func (t *Triangle) Area() float64 {
	return (t.bottom * t.height) / 2
}

func (s *Square) Area() float64 {
	return s.length * s.length
}

func main() {
	t := Triangle{6, 8}
	s := Square{8}
	shapes := []Shape{&t, &s}
	for _, v := range shapes {
		fmt.Printf("fileds: %.2f, Shape Area: %.2f \n", v, v.Area())
	}
}

/*
fileds: &{6.00 8.00}, Shape Area: 24.00
fileds: &{8.00}, Shape Area: 64.00
*/
