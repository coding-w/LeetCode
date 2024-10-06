package main

import "sort"

func main() {
	res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	for _, v := range res {
		for _, vv := range v {
			print(vv, " ")
		}
		println()
	}

}

// https://leetcode.cn/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 1 {
		if candidates[0] == target {
			return [][]int{candidates}
		} else {
			return [][]int{}
		}
	}
	sort.Ints(candidates)
	temp := make([]int, 0)
	res := make([][]int, 0)
	recursion(0, 0, candidates, target, &res, &temp)

	return res

}

func recursion(index, sum int, candidates []int, target int, res *[][]int, temp *[]int) {
	if sum == target {
		tempCopy := make([]int, len(*temp))
		copy(tempCopy, *temp)
		*res = append(*res, tempCopy)
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		*temp = append(*temp, candidates[i])
		recursion(i+1, sum+candidates[i], candidates, target, res, temp)
		*temp = (*temp)[:len(*temp)-1]
	}
}
