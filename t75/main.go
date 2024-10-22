package main

import "sort"

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

// https://leetcode.cn/problems/sort-colors/description/
func sortColors(nums []int) {
	sort.Ints(nums)
}
