package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]

}
