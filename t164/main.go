package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumGap([]int{1, 3, 100}))
}

// https://leetcode.cn/problems/maximum-gap/
// tips：桶排序..
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	num := 0
	for i := 0; i < len(nums)-1; i++ {
		num = max(num, nums[i+1]-nums[i])
	}
	return num
}
