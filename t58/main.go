package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

// https://leetcode.cn/problems/length-of-last-word/description/
func lengthOfLastWord(s string) int {
	n := len(s)
	if n == 1 {
		if s[0] == ' ' {
			return 0
		} else {
			return 1
		}
	}
	res := 0
	idx := n - 1
	for idx >= 0 {
		if s[idx] != ' ' {
			break
		}
		idx--
	}
	for i := idx; i >= 0; i-- {
		if s[i] != ' ' {
			res++
		} else {
			break
		}
	}
	return res
}
