package main

func main() {

}

// https://leetcode.cn/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
	}
	res := -10001
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}
