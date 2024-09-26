package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var numMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var res []string
	var build func(int, string)
	build = func(index int, cur string) {
		if index == len(digits) {
			res = append(res, cur)
			return
		}
		for _, v := range numMap[string(digits[index])] {
			build(index+1, cur+string(v))
		}
	}
	build(0, "")

	return res
}
