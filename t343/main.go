package main

func main() {

}

// https://leetcode.cn/problems/integer-break/
// dp
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], (i-j)*j, dp[i-j]*j)
		}
	}
	return dp[n]
}
