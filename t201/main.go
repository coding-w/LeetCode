package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(1, 1000))
}

// https://leetcode.cn/problems/bitwise-and-of-numbers-range/
func rangeBitwiseAnd(left int, right int) int {
	i := 0
	for left != right {
		left >>= 1
		right >>= 1
		i++
	}
	return left << i
}
