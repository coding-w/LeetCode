package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{2, 3, 1}))
}

// https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	var dfs func()
	var ans [][]int
	var ns []int
	var state = make(map[int]bool)
	n := len(nums)
	dfs = func() {
		if len(ns) == n {
			tmp := make([]int, n)
			copy(tmp, ns)
			ans = append(ans, tmp)
			return
		}
		for i, num := range nums {
			if state[i] {
				continue
			}
			ns = append(ns, num)
			state[i] = true
			dfs()
			state[i] = false
			ns = ns[:len(ns)-1]
		}
	}
	dfs()
	return ans
}
