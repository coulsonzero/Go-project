package main

import (
	"testing"
)

/*
$ go test -bench .
goos: darwin
goarch: arm64
pkg: go-study/fuzzing/benchmark/struct-benchmark
BenchmarkValueReceiver-8        1000000000               0.3163 ns/op
BenchmarkPointerReceiver-8      1000000000               0.3128 ns/op
PASS
ok      go-study/fuzzing/benchmark/struct-benchmark     0.797s
*/

type Struct struct {
	value int
}

// Value Receiver
func (s Struct) plusOne() int {
	return s.value + 1
}

// Pointer Receiver
func (s *Struct) plusOnePointer() int {
	return s.value + 1
}

// Benchmark
func BenchmarkValueReceiver(b *testing.B) {
	s := Struct{value: 1}
	for i := 0; i < b.N; i++ {
		s.plusOne()
	}
}

func BenchmarkPointerReceiver(b *testing.B) {
	s := &Struct{value: 1}
	for i := 0; i < b.N; i++ {
		s.plusOnePointer()
	}
}
