package main

import "fmt"

func main() {
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))
}

// https://leetcode.cn/problems/number-of-ways-to-split-array/?envType=daily-question&envId=2025-01-13
// 每日一题
func waysToSplitArray(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	res := 0
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[n-1]-nums[i-1] {
			res++
		}
	}
	return res
}
