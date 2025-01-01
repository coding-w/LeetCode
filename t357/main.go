package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(8))
}

// https://leetcode.cn/problems/count-numbers-with-unique-digits/
// 排列组合
func countNumbersWithUniqueDigits(n int) int {
	res := 1
	product := 9
	for i := 1; i <= n; i++ {
		res += product
		product *= 10 - i
	}
	return res
}
