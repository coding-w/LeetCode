package main

import (
	"fmt"
)

func main() {
	//fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

// https://leetcode.cn/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	dp := make([]bool, l+1)
	dp[0] = true
	for i := 1; i <= l; i++ {
		for _, v := range wordDict {
			index := i - len(v)
			if index >= 0 && dp[index] && v == s[index:i] {
				dp[i] = true
			}
		}
	}
	return dp[l]
}
