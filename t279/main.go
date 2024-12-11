package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(18))
}

// https://leetcode.cn/problems/perfect-squares/
func numSquares(n int) int {
	sqrt := int(math.Sqrt(float64(n)))
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = 10001
	}
	dp[0] = 0
	for i := 1; i <= sqrt; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}
