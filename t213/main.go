package main

import "fmt"

func main() {

	fmt.Println(rob([]int{1, 2, 3}))
}

// https://leetcode.cn/problems/house-robber-ii/description/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := len(nums)
	dp1 := make([]int, l+1)
	dp2 := make([]int, l+1)
	dp1[1] = nums[0]
	dp2[2] = nums[1]
	for i := 1; i < l-1; i++ {
		dp1[i+1] = max(nums[i]+dp1[i-1], dp1[i])
	}
	for i := 2; i < l; i++ {
		dp2[i+1] = max(nums[i]+dp2[i-1], dp2[i])
	}
	return max(dp1[l-1], dp2[l])
}
