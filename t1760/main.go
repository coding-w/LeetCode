package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumSize([]int{2, 4, 8, 2}, 4))
}

// https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/
// 每日一题
func minimumSize(nums []int, maxOperations int) int {
	left, right := 0, slices.Max(nums)
	for left+1 < right {
		mid := left + (right-left)>>1
		sum := 0
		for _, num := range nums {
			sum += (num - 1) / mid
		}
		if sum <= maxOperations {
			right = mid
		} else {
			left = mid
		}
	}

	return right
}
