package main

import "sort"

func main() {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			print(" ", res[i][j])
		}
		print("\n")
	}
}

// https://leetcode.cn/problems/combination-sum/description/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 1 {
		if candidates[0] == target {
			res = append(res, []int{candidates[0]})
		} else {
			return res
		}
	}
	sort.Ints(candidates)
	temp := make([]int, 0)
	recursion(0, 0, candidates, target, &res, &temp)

	return res
}
func recursion(i, sum int, candidates []int, target int, res *[][]int, temp *[]int) {
	if sum > target || i >= len(candidates) {
		return
	}
	if sum == target {
		tt := make([]int, len(*temp))
		copy(tt, *temp)
		*res = append(*res, tt)
		return
	}
	*temp = append(*temp, candidates[i])
	recursion(i, sum+candidates[i], candidates, target, res, temp)
	*temp = (*temp)[:len(*temp)-1]
	recursion(i+1, sum, candidates, target, res, temp)

}
