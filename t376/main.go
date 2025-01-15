package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

// https://leetcode.cn/problems/wiggle-subsequence/
// 贪心
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	res := 1
	preNum := 0
	for i := 1; i < n; i++ {
		curNum := nums[i] - nums[i-1]
		if preNum >= 0 && curNum < 0 || preNum <= 0 && curNum > 0 {
			preNum = curNum
			res++
		}
	}
	return res
}
