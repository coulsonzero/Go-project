package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// |x|
	abs := math.Abs(123.2)
	fmt.Println(abs)

	// 随机数
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))

	// 取整
	ceil := math.Ceil(12.6)
	floor := math.Floor(12.6)
	fmt.Println(ceil, floor)

	r := rand.New(rand.NewSource(99))
	println(r.Intn(10))

	// 打乱随机顺序
	nums := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	fmt.Println(nums)

	// 生成随机数组
	fmt.Println(rand.Perm(4))

}
