package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return []int{-1, -1}
	}
	for right < len(nums) && nums[right] == target {
		right++
	}
	return []int{left, right - 1}
}
