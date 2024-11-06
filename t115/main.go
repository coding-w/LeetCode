package main

import "fmt"

func main() {
	//fmt.Println(numDistinct("babgbag", "bag"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

// https://leetcode.cn/problems/distinct-subsequences/
// tips: 动态规划
// s[i] == t[j], dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
// s[i] != t[j], dp[i][j] = dp[i][j-1]
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
