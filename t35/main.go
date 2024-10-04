package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}

// https://leetcode.cn/problems/search-insert-position/description/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if nums[0] >= target {
		return 0
	}
	return left
}
