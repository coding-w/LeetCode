package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))
}

// https://leetcode.cn/problems/maximum-number-of-coins-you-can-get/?envType=daily-question&envId=2025-01-22
func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)
	res := 0
	for n >= 3 {
		res += piles[n-2]
		n -= 2
		piles = piles[:n]
		piles = piles[1:]
		n--
	}
	return res
}
