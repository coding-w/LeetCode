package main

import "fmt"

func main() {
	fmt.Println(isValid("{}{[]}"))
}

// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
