package main

func appendCharacters(s string, t string) int {
	n, m := len(s), len(t)
	j := 0
	for i := 0; i < n && j < m; i++ {
		if s[i] == t[j] {
			j++
		}
	}
	return m - j
}

func main() {
	println(appendCharacters("coaching", "coding")) // 4, 末尾追加字符串 "ding"
	println(appendCharacters("abcde", "a"))         // 0, 末尾无需追加字符串
	println(appendCharacters("z", "abcde"))         // 5, 末尾追加字符串"abcde"
	println(appendCharacters("vrykt", "rkge"))      // 2, 末尾追加字符串 "ge"
}
