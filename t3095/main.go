package main

import (
	"fmt"
)

func main() {
	fmt.Println(16 | 20 | 32)
	fmt.Println(52 ^ 16)
	fmt.Println(11 ^ 1)
	fmt.Println(minimumSubarrayLength([]int{16, 1, 2, 20, 32}, 45))
}

// https://leetcode.cn/problems/shortest-subarray-with-or-at-least-k-i/description/?envType=daily-question&envId=2025-01-16
// 每日一题
func minimumSubarrayLength(nums []int, k int) int {
	res := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		cur := 0
		for j := i; j < len(nums); j++ {
			cur |= nums[j]
			if cur >= k && res > j-i+1 {
				res = j - i + 1
			}
		}
	}
	if res == len(nums)+1 {
		return -1
	}
	return res
}
