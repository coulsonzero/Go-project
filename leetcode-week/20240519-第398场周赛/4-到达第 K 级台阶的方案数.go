package main


func waysToReachStair(k int) (ans int) {
	type pair struct {
		x, y int
		pre  bool
	}
	memo := map[pair]int{}
	var f func(int, int, bool) int
	f = func(i, j int, pre bool) (res int) {
		if i > k+1 {
			return
		}
		p := pair{i, j, pre}
		if v, ok := memo[p]; ok {
			return v
		}

		if i == k {
			res = 1
		}
		r := f(i+1<<j, j+1, false)
		res += r
		if !pre && i > 0 {
			r2 := f(i-1, j, true)
			res += r2
		}
		memo[p] = res
		return
	}
	ans = f(1, 0, false)
	return
}