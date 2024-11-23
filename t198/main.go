package main

func main() {

}

// https://leetcode.cn/problems/house-robber/
func rob(nums []int) int {
	l := len(nums)
	dp := make([]int, l+1)
	dp[1] = nums[0]
	for i := 1; i < l; i++ {
		dp[i+1] = max(nums[i]+dp[i-1], dp[i])
	}
	return dp[l]
}
