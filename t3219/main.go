package main

import (
	"slices"
)

func main() {

}

// https://leetcode.cn/problems/minimum-cost-for-cutting-cake-i/
// 每日一题
// 贪心， 先切较大的开销
func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	slices.SortFunc(horizontalCut, func(a, b int) int {
		return b - a
	})
	slices.SortFunc(verticalCut, func(a, b int) int {
		return b - a
	})
	i, j := 0, 0
	res := 0
	for i < m-1 && j < n-1 {
		if horizontalCut[i] < verticalCut[j] {
			res += verticalCut[j] * (i + 1)
			j++
		} else {
			res += horizontalCut[i] * (j + 1)
			i++
		}
	}
	for i < m-1 {
		res += horizontalCut[i] * (j + 1)
		i++
	}
	for j < n-1 {
		res += verticalCut[j] * (i + 1)
		j++
	}
	return int64(res)
}
