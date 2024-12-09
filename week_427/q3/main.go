package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySum([]int{1, 2}, 1))

}

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	var res int64
	if k == 0 {
		res = int64(nums[0])
	} else {
		res = math.MinInt64
	}
	dp := make([]int64, n)
	dp[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + int64(nums[i])
	}
	kk := k
	for n >= k {
		res = max(res, dp[k-1])
		for i := k; i < n; i++ {
			res = max(res, dp[i]-dp[i-k])
		}
		k += kk
	}

	return res
}
