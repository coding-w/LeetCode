package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
	//s := "123*"
	//_, i := beforeNum(s, 3)
	//fmt.Println(s[:i])
	//fmt.Println(strconv.Atoi("+123"))
}

// https://leetcode.cn/problems/different-ways-to-add-parentheses/description/
func diffWaysToCompute(expression string) []int {
	hmap := make(map[string][]int)
	var dfs func(s string) []int
	dfs = func(s string) []int {
		if v, ok := hmap[s]; ok {
			return v
		}
		var res []int
		// 如果字符串中没有运算符，直接返回解析的数字
		if !strings.ContainsAny(s, "+-*") {
			num, _ := strconv.Atoi(s)
			return []int{num}
		}

		// 遍历表达式，按运算符拆分
		for i := 0; i < len(s); i++ {
			if s[i] == '+' || s[i] == '-' || s[i] == '*' {
				// 递归计算左右子表达式
				left := dfs(s[:i])
				right := dfs(s[i+1:])

				// 根据运算符计算结果
				for _, l := range left {
					for _, r := range right {
						var resVal int
						switch s[i] {
						case '+':
							resVal = l + r
						case '-':
							resVal = l - r
						case '*':
							resVal = l * r
						}
						res = append(res, resVal)
					}
				}
			}
		}
		hmap[s] = res
		return res
	}
	return dfs(expression)
}
