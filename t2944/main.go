package main

func main() {

}

func minimumCoins(prices []int) int {
	n := len(prices)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + prices[i-1]
		for j := i - 1; j*2 >= i; j-- {
			dp[i] = min(dp[i], dp[j-1]+prices[j-1])
		}
	}
	return dp[n]
}
