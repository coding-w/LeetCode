package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}

// https://leetcode.cn/problems/increasing-triplet-subsequence/
// 注重思路！！！！！！
// a,b常量
func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt32, math.MaxInt32
	for _, n := range nums {
		if n <= a {
			a = n
		} else if n <= b {
			b = n
		} else {
			return true
		}
	}
	return false
}
