package main

func main() {

}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit(prices []int) int {
	minl, maxr, maxTarget := prices[0], prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minl {
			minl = prices[i]
			maxr = prices[i]
		}
		if prices[i] > maxr {
			maxr = prices[i]
		}
		if maxr-minl > 0 {
			maxTarget += maxr - minl
			minl = prices[i]
			maxr = prices[i]
		}
	}
	if maxTarget > 0 {
		return maxTarget
	}
	return 0
}
