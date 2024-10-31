package main

import "fmt"

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}

// https://leetcode.cn/problems/interleaving-string/description/
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	m, n := len(s1), len(s2)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 1; i <= m; i++ {
		if s1[i-1] != s3[i-1] {
			break
		}
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		if s2[i-1] != s3[i-1] {
			break
		}
		dp[0][i] = true
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if !dp[i-1][j] && !dp[i][j-1] {
				continue
			}
			if (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1]) {
				dp[i][j] = true
			}
		}
	}
	return dp[m][n]
}
