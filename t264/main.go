package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

// https://leetcode.cn/problems/ugly-number-ii/
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	i1, i2, i3 := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(2*dp[i1], 3*dp[i2], 5*dp[i3])
		if dp[i] == 2*dp[i1] {
			i1++
		}
		if dp[i] == 3*dp[i2] {
			i2++
		}
		if dp[i] == 5*dp[i3] {
			i3++
		}
	}
	return dp[n-1]
}
