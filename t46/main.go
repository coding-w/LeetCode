package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{2, 3, 1}))
}

// https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	ll := len(nums)
	var res [][]int
	if ll == 0 {
		return [][]int{}
	}
	if ll == 1 {
		return [][]int{nums}
	}
	res = [][]int{}
	var dfs func(nums []int, temp []int)
	dfs = func(nums []int, temp []int) {
		if len(temp) == ll {
			res = append(res, append([]int(nil), temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			aa := nums[i]
			dfs(append(append([]int(nil), nums[:i]...), nums[i+1:]...), append(temp, aa))
		}
	}
	dfs(nums, []int{})
	return res

}
