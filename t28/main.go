package main

import "fmt"

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(needle) == 1 {
				return i
			}
			for j := 0; j < len(needle); j++ {
				if i+j >= len(haystack) {
					return -1
				}
				if haystack[i+j] != needle[j] {
					break
				}
				if j == len(needle)-1 {
					return i
				}
			}
		}
	}
	return -1
}
