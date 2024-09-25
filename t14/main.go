package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"Qwe", "Qwer", "Q"}))
}

// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	if len(strs[0]) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	index := 0
	res := ""
	temp := ""
	for index < len(strs[0]) {
		temp = strs[0][0 : index+1]
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) == 0 {
				return ""
			}
			if len(strs[i]) <= index || strs[i][0:index+1] != temp {
				return res
			}
		}
		res = temp
		index++

	}
	return res
}
