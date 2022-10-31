package main

import "testing"

// 压力测试
// go test -bench . -benchmem -gcflags "-N -l"

// BenchmarkArray-8          585766              3016 ns/op               0 B/op          0 allocs/op
// BenchmarkSlice-8          403167              2846 ns/op            8192 B/op          1 allocs/op
// PASS
// ok      go-test/testing 3.861s

// go test array_slice_test.go -test.bench=".*" .

// BenchmarkArray-8         2010756               560.6 ns/op
// BenchmarkSlice-8         3628813               329.8 ns/op
// PASS
// ok      command-line-arguments  3.591s

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}
