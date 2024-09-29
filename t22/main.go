package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// https://leetcode.cn/problems/generate-parentheses/description/
func generateParenthesis(n int) []string {
	l, r := n, n
	res := make([]string, 0)
	var recursion func(int, int, string)
	recursion = func(l, r int, s string) {
		if l == 0 && r == 0 {
			res = append(res, s)
		}
		if l > 0 {
			recursion(l-1, r, s+"(")
		}
		if r > l {
			recursion(l, r-1, s+")")
		}
	}
	recursion(l, r, "")
	return res
}
