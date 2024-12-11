package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/h-index/
// 数组中有h个不小于h的值，求最大的h
func hIndex(citations []int) int {

	n := len(citations)
	sort.Ints(citations)
	for i := n - 1; i >= 0; i-- {
		if n-i >= citations[i] {
			return max(n-i-1, citations[i])
		}
	}
	return n
}
