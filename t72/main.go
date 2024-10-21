package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

// https://leetcode.cn/problems/edit-distance/description/
// tips: 由少到多递推
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
