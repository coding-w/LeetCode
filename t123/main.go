package main

func main() {

}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit(prices []int) int {

	// 对于任意一天考虑四个变量:
	// buy1: 在该天第一次买入股票可获得的最大收益
	// sell1: 在该天第一次卖出股票可获得的最大收益
	// buy2: 在该天第二次买入股票可获得的最大收益
	// sell2: 在该天第二次卖出股票可获得的最大收益
	// 分别对四个变量进行相应的更新, 最后secSell就是最大
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for _, price := range prices {
		buy1 = max(buy1, -price)
		sell1 = max(sell1, buy1+price)
		buy2 = max(buy2, sell1-price)
		sell2 = max(sell2, buy2+price)
	}
	return sell2
}
