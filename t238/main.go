package main

import (
	"fmt"
)

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

// https://leetcode.cn/problems/product-of-array-except-self/
// 前缀集 and 后缀集
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	item := 1
	for i := 0; i < len(nums); i++ {
		res[i] = item
		item *= nums[i]
	}
	item = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= item
		item *= nums[i]
	}
	return res
}
