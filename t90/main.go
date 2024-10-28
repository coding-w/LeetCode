package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(subsetsWithDup([]int{1, 2, 2, 2}))

}

// https://leetcode.cn/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}
	start, end := 0, 0
	// 遍历每个数字
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			start = end + 1
		} else {
			start = 0
		}
		end = len(res) - 1
		for j := start; j <= end; j++ {
			newSubset := append([]int{}, res[j]...)
			newSubset = append(newSubset, nums[i])
			res = append(res, newSubset)
		}
	}
	return res
}
