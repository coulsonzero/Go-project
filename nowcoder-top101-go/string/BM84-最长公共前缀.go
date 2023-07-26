package main

func longestCommonPrefix(strs []string) string {
	// write code here
	if len(strs) == 0 {
		return ""
	}
	index := 0
	for {
		if len(strs[0]) == index {
			return strs[0][:index]
		}
		prefix := strs[0][index]
		for _, str := range strs {
			if len(str) == index {
				return strs[0][:index]
			}
			if str[index] != prefix {
				return strs[0][:index]
			}
		}
		index++
	}
}
