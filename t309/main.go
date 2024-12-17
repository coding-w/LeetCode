package main

func main() {

}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	// 0 买入
	// 1 卖出
	dp[0][0] = -prices[0]
	if len(prices) > 1 {
		dp[1][0] = max(dp[0][0], -prices[1])
		dp[1][1] = max(dp[0][1], dp[0][0]+prices[1])
	}

	for i := 2; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]

}
