package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 1}))

}

// https://leetcode.cn/problems/permutations-ii/description/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
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
		} else {
			for i := 0; i < n; i++ {
				if !state[i] {
					ns = append(ns, nums[i])
					state[i] = true
					dfs()
					state[i] = false
					ns = ns[:len(ns)-1]
					temp := nums[i]
					for i < n && nums[i] == temp {
						i++
					}
					i--
				}
			}
		}
	}
	dfs()
	return ans
}
