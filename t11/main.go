package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maxArea(nums))
}

// https://leetcode.cn/problems/container-with-most-water/description/
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		if height[l] < height[r] {
			res = compare(res, (r-l)*height[l])
			l++
		} else {
			res = compare(res, (r-l)*height[r])
			r--
		}
	}
	return res
}

func compare(a, b int) int {
	if a > b {
		return a
	}
	return b
}
