package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
}

// https://leetcode.cn/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	res, tmp := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			tmp++
		} else {
			res = max(res, tmp)
			tmp = 1
		}
	}
	return max(res, tmp)
}
