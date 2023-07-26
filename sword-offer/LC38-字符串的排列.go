package main

func permutation(s string) []string {
	var res []string // 返回值列表
	Hmap := make(map[byte]int)
	ele := "" // 待构造的字符串
	for i := 0; i < len(s); i++ {
		Hmap[s[i]]++ // 计数
	}
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			res = append(res, ele) // 字符串构造完毕 添加进返回值列表
			return
		}
		for k, _ := range Hmap {
			if Hmap[k] != 0 { // 次数不为0说明可用
				ele += string(k)
				Hmap[k]--
				dfs(start + 1)         // 从新的点继续构造字符串
				ele = ele[:len(ele)-1] // 回溯
				Hmap[k]++
			}
		}
	}
	dfs(0)
	return res
}
