package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

// https://leetcode.cn/problems/jump-game-ii/description/
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}
	return dp[n-1]
}
