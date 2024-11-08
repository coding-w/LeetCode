package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("0p"))

}

// https://leetcode.cn/problems/valid-palindrome/
func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^0-9a-zA-Z]+")
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
