package main

import "fmt"

func main() {

	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))

}

// https://leetcode.cn/problems/max-sum-of-rectangle-no-larger-than-k/
func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
			if matrix[i-1][j-1] == k || dp[i][j] == k {
				return k
			}
		}
	}
	res := -100001
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			for x := i; x <= m; x++ {
				for y := j; y <= n; y++ {
					cur := dp[x+1][y+1] - dp[i][y+1] - dp[x+1][j] + dp[i][j]
					if cur == k {
						return k
					}
					if cur < k {
						res = max(res, cur)
					}
				}

			}
		}
	}
	return res
}
