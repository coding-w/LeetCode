package main

import "math"

func main() {

}

// https://leetcode.cn/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	dpMax, dpMin, res := 1, 1, math.MinInt32
	for _, num := range nums {
		tmp := dpMax
		dpMax = max(dpMax*num, dpMin*num, num)
		dpMin = min(dpMin*num, tmp*num, num)
		res = max(res, dpMax)
	}
	return res
}
