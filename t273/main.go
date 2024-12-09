package main

import "strings"

func main() {

}

var lessThan20 = []string{
	"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
}

var tens = []string{
	"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
}

var thousands = []string{
	"", "Thousand", "Million", "Billion",
}

// https://leetcode.cn/problems/integer-to-english-words/description/
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	result := ""
	i := 0
	// 分别处理每一段的数字
	for num > 0 {
		if num%1000 != 0 {
			result = helper(num%1000) + thousands[i] + " " + result
		}
		num /= 1000
		i++
	}
	return strings.TrimSpace(result) // 去除多余的空格
}

// 处理三位数以内的数字
func helper(num int) string {
	if num == 0 {
		return ""
	}
	if num < 20 {
		return lessThan20[num] + " "
	}
	if num < 100 {
		return tens[num/10] + " " + helper(num%10)
	}
	return lessThan20[num/100] + " Hundred " + helper(num%100)
}
