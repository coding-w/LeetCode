package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

// https://leetcode.cn/problems/longest-palindromic-substring/
// tips 中心扩展
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	res := ""
	for i := 0; i < len(s); i++ {
		// 该字符为奇数中心
		odd := center(s, i, i)

		// 该字符以偶数为中心
		even := center(s, i, i+1)

		if len(odd) > len(res) {
			res = odd
		}
		if len(even) > len(res) {
			res = even
		}
	}
	return res
}

func center(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}
