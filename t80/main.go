package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
