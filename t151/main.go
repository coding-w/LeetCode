package main

import "strings"

func main() {

}

// https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	sb := strings.Builder{}
	res := ""
	for _, ch := range s {
		if ch != ' ' {
			sb.WriteRune(ch)
		} else {
			if sb.Len() > 0 {
				res = sb.String() + " " + res
				sb.Reset()
			}
		}
	}
	if sb.Len() > 0 {
		res = sb.String() + " " + res
	}
	return res[:len(res)-1]
}
