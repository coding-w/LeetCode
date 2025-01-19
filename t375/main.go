package main

import "math"

func main() {

}

// https://leetcode.cn/problems/guess-number-higher-or-lower-ii/description/
// https://leetcode.cn/problems/guess-number-higher-or-lower-ii/solutions/83395/dong-tai-gui-hua-c-you-tu-jie-by-zhang-xiao-tong-2/
// 顶级题解
func getMoneyAmount(n int) int {
	if n == 1 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
		dp[i][i] = 0
	}
	for j := 2; j <= n; j++ {
		for i := j - 1; i >= 1; i-- {
			for k := i + 1; k <= j-1; k++ {
				dp[i][j] = min(k+max(dp[k+1][j], dp[i][k-1]), dp[i][j])
			}
			dp[i][j] = min(dp[i][j], j+dp[i][j-1], i+dp[i+1][j])
		}
	}
	return dp[1][n]
}
