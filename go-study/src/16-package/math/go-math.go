package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Perm(10))

}

func demo() {
	// |x|
	abs := math.Abs(123.2)
	fmt.Println(abs)

	// 随机数 (1, 10)
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(10))

	// 取整
	ceil := math.Ceil(12.6)   // 13
	floor := math.Floor(12.6) // 12
	fmt.Println(ceil, floor)

	r := rand.New(rand.NewSource(99))
	println(r.Intn(10))

	// 打乱随机顺序
	nums := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	fmt.Println(nums)

	// 生成随机数组
	fmt.Println(rand.Perm(4))

	// 幂
	println(math.Pow(10, 2))
	println(math.Pow10(2))
}
