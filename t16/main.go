package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
}

// https://leetcode.cn/problems/3sum-closest/description/
func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	var res int
	var resTemp = math.MaxInt32
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		left, right := i+1, n-1
		// 简化成两数之和
		for left < right {
			temp := nums[i] + nums[left] + nums[right]
			if abs(target-temp) < resTemp {
				resTemp = abs(target - temp)
				res = temp
			}
			if target-temp == 0 {
				return temp
			} else if target > temp {
				// 小于 0
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
