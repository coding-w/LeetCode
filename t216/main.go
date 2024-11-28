package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 9))
}

// https://leetcode.cn/problems/combination-sum-iii/description/
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	if k > n {
		return res
	}
	var backtrack func([]int, int, int)
	backtrack = func(tmp []int, start int, sum int) {
		if len(tmp) == k && sum == n {
			res = append(res, append([]int{}, tmp...))
		}
		for i := start; i <= 9; i++ {
			if sum+i > n {
				continue
			}
			sum += i
			tmp = append(tmp, i)
			backtrack(tmp, i+1, sum)
			tmp = tmp[:len(tmp)-1]
			sum -= i
		}
	}
	backtrack([]int{}, 1, 0)
	return res
}
