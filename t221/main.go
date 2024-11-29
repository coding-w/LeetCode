package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}

// https://leetcode.cn/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxN := 0
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		dp[i][0] = int(matrix[i][0] - '0')
		maxN = max(dp[i][0], maxN)
	}
	for i := 0; i < len(matrix[0]); i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		maxN = max(dp[0][i], maxN)
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			pre := min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			if pre == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + pre
			}
			maxN = max(maxN, dp[i][j])
		}
	}
	return maxN * maxN
}
