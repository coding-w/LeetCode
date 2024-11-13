package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

// https://leetcode.cn/problems/word-break-ii/
func wordBreak(s string, wordDict []string) []string {
	res := make([]string, 0)
	l := len(s)
	var check func(string) bool
	check = func(s string) bool {
		for _, w := range wordDict {
			if s == w {
				return true
			}
		}
		return false
	}
	var dfs func(int, []string)
	dfs = func(index int, tmp []string) {
		if index == l {
			res = append(res, strings.Join(tmp, " "))
			return
		}
		for i := index; i <= l; i++ {
			if check(s[index:i]) {
				tmp = append(tmp, s[index:i])
				dfs(i, tmp)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0, []string{})
	return res
}
