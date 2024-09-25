package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, 2}))
}

// https://leetcode.cn/problems/3sum/description/
func threeSum(nums []int) [][]int {
	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		// nums[i] == nums[i-1] 去除相同解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		// 简化成两数之和
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				// 小于 0
				left++
			} else {
				right--
			}
		}
	}
	return res
}
