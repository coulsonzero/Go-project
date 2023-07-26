package main

import (
	"math/rand"
	"testing"
	"time"
)

/**
 * $ go test -bench .
 * goos: darwin
 * goarch: arm64
 * pkg: math
 * BenchmarkGenerate1000-8            42018             26012 ns/op
 * BenchmarkGenerate10000-8            6951            170476 ns/op
 * BenchmarkGenerate100000-8            747           1592522 ns/op
 * BenchmarkGenerate1000000-8            68          16441272 ns/op
 * PASS
 * ok      math 5.666s
 */

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i)
	}
}

func BenchmarkGenerate1000(b *testing.B)    { benchmarkGenerate(1000, b) }
func BenchmarkGenerate10000(b *testing.B)   { benchmarkGenerate(10000, b) }
func BenchmarkGenerate100000(b *testing.B)  { benchmarkGenerate(100000, b) }
func BenchmarkGenerate1000000(b *testing.B) { benchmarkGenerate(1000000, b) }
