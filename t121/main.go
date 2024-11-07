package main

func main() {

}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
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
		if maxr-minl > maxTarget {
			maxTarget = maxr - minl
		}
	}
	if maxTarget > 0 {
		return maxTarget
	}
	return 0
}
