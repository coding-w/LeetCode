package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3749))
}

// https://leetcode.cn/problems/integer-to-roman/description/
func intToRoman(num int) string {
	rMap := []struct {
		key   int
		value string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := strings.Builder{}
	for _, item := range rMap {
		for num >= item.key {
			res.WriteString(item.value)
			num -= item.key
		}
	}
	return res.String()
}
