package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isNumber("-8115e957"))
}

func isNumber(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	i := 0
	if s[i] == '+' || s[i] == '-' {
		i++
	}
	isValid := false
	for i < n && unicode.IsDigit(rune(s[i])) {
		i++
		isValid = true
	}

	if i < n && s[i] == '.' {
		i++
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isValid = true
		}
	}

	if !isValid {
		return isValid
	}
	// e的情况
	if i < n && (s[i] == 'e' || s[i] == 'E') {
		i++
		if i == n {
			// e后面必须跟着一个整数
			return false
		}

		if s[i] == '+' || s[i] == '-' {
			i++
		}

		isValid = false
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isValid = true
		}

		if !isValid {
			return isValid
		}
	}

	return i == n
}
