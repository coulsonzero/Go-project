package main

import "strings"

func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int)

	for _, word := range words {
		// temp := ""
		temp := strings.Builder{}
		for _, v := range word {
			// temp += morse[v-'a']
			temp.WriteString(morse[v-'a'])
		}
		// m[temp]++
		m[temp.String()]++
		// fmt.Println(temp, m)
	}

	return len(m)
}

func main() {
	println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})) // 2
	println(uniqueMorseRepresentations([]string{"a"}))                        // 1
}
