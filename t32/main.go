package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()((())"))
}

// https://leetcode.cn/problems/longest-valid-parentheses/
// 大佬评论
// 用栈模拟一遍，将所有无法匹配的括号的位置全部置1
// 例如: "()(()"的mark为[0, 0, 1, 0, 0]
// 再例如: ")()((())"的mark为[1, 0, 0, 1, 0, 0, 0, 0]
// 经过这样的处理后, 此题就变成了寻找最长的连续的0的长度
// 茅塞顿开
func longestValidParentheses(s string) int {
	tmp := make([]int, len(s))
	index := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			index = append(index, i)
		}
		if s[i] == ')' && len(index) > 0 {
			j := index[len(index)-1]
			index = index[:len(index)-1]
			tmp[i] = 1
			tmp[j] = 1
		}
	}
	res, tt := 0, 0
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 1 {
			tt++
		} else {
			res = max(res, tt)
			tt = 0
		}
	}
	res = max(res, tt)
	return res
}
