package main

/**
 * Go中可以抛出一个panic的异常，
 * 然后在defer中通过recover捕获这个异常，然后正常处理。
 */

func temperature(t float64) (res string) {
	// write code here
	defer func() {
		if err := recover(); err != nil {
			res = "体温异常"
		}
	}()

	if t > 37.5 {
		panic("体温异常")
	}
	return res
}
