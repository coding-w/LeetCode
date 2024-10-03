package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
}

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		// 找到目标值
		if nums[mid] == target {
			return mid
		}
		// 判断升序列
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				// 在升序序列中
				r = mid - 1
			} else {
				// 在非升序序列中
				l = mid + 1
			}
		} else {
			// 不在升序序列中
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
