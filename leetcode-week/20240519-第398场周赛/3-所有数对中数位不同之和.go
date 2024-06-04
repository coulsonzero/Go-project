package main

import "strconv"


func sumDigitDifferences(a []int) int64 {
	ans64 := func() (ans int) {
		m := len(strconv.Itoa(a[0]))
		cnt := make([][]int, m)
		for i := range cnt {
			cnt[i] = make([]int, 10)
		}
		for _, v := range a {
			for i, b := range strconv.Itoa(v) {
				bb := int(b-'0')
				for j, c := range cnt[i] {
					if j != bb {
						ans += c
					}
				}
				cnt[i][bb]++
			}
		}
		return
	}()
	return int64(ans64)
}