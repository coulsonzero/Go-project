package main

import "strings"

func interpret(command string) string {
	res := strings.ReplaceAll(command, "(al)", "al")
	return strings.ReplaceAll(res, "()", "o")
}

func interpret2(command string) string {
	var res strings.Builder
	for i, c := range command {
		if c == 'G' {
			res.WriteByte('G')
		} else if c == '(' {
			if command[i+1] == ')' {
				res.WriteByte('o')
			} else {
				res.WriteString("al")
			}
		}
	}
	return res.String()
}

func interpret3(command string) string {
	res := ""
	start, end := 0, 0
	for i := 0; i < len(command); i++ {
		if command[i] == '(' {
			start = i
		} else if command[i] == ')' {
			end = i
			if command[start:end] == "(" {
				res += "o"
			}
		} else {
			res += string(command[i])
		}
	}
	return res
}
