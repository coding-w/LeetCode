package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/maximum-consecutive-floors-without-special-floors/?envType=daily-question&envId=2025-01-06
// 每日一题
func maxConsecutive(bottom int, top int, special []int) int {
	res := special[0] - bottom
	sort.Ints(special)
	for i := 1; i < len(special)-1; i++ {
		res = max(res, special[i]-special[i-1]-1)
	}
	res = max(res, top-special[len(special)-1])
	return res
}
