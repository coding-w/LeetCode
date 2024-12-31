package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumCost(3, 2, []int{1, 3}, []int{5}))
}

// https://leetcode.cn/problems/minimum-cost-for-cutting-cake-i/
// 贪心， 先切较大的开销
func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
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
	return res
}
