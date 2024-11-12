package main

import "fmt"

func main() {
	fmt.Println(2 ^ 3 ^ 2)
}

// https://leetcode.cn/problems/single-number/
// tips: 异或
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
