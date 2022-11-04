package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	sequence := "aaabaaaabaaabaaaabaaaabaaaabaaaaba"
	word := "aaaba"
	// count := 0
	// temp := []int{}
	// for i := 0; i <= len(sequence)-len(word); {
	// 	if word == sequence[i:len(word)+i] {
	// 		count++
	// 		i += len(word)
	// 	} else {
	// 		i++;
	// 		if count != 0 {
	// 			temp = append(temp, count)
	// 			count = 0
	// 		}
	// 	}
	//
	// }
	// fmt.Println(count)
	// fmt.Println(temp)
	fmt.Println(len(sequence) / len(word))
	fmt.Println(strings.Repeat(word, 1))
	fmt.Println(strings.Index(sequence, strings.Repeat(word, 1)))

	count := 0
	for i := len(sequence) / len(word); i > 0; i-- {
		if strings.Index(sequence, strings.Repeat(word, i)) != -1 {
			count = i
			fmt.Println(count)
		}
	}
	fmt.Println(count)
}

func demo() {
	// |x|
	abs := math.Abs(123.2)
	fmt.Println(abs)

	// 随机数
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))

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
}
