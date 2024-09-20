package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	res := 0
	ss := ""
	for _, v := range s {
		index := strings.Index(ss, string(v))
		if index >= 0 {
			ss += string(v)
			ss = ss[index+1:]
		} else {
			ss += string(v)
		}
		if len(ss) > res {
			res = len(ss)
		}
	}
	return res
}
