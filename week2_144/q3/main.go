package main

import "fmt"

func main() {
	//fmt.Println(maxRemoval([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}))
	fmt.Println(maxRemoval([]int{1, 1, 1, 1}, [][]int{{1, 3}, {0, 2}, {1, 2}, {1, 3}}))
}

func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)
	m := len(queries)

	// 用于检查剩余的 queries 是否能够将 nums 变为零数组
	check := func(active []bool) bool {
		// 差分数组模拟
		diff := make([]int, n)
		for i := 0; i < m; i++ {
			if !active[i] {
				l, r := queries[i][0], queries[i][1]
				diff[l]++
				if r+1 < n {
					diff[r+1]--
				}
			}
		}

		// 使用差分数组检查 nums 是否可以归零
		curr := 0
		for i := 0; i < n; i++ {
			curr += diff[i]
			if nums[i]-curr > 0 {
				return false
			}
		}
		return true
	}

	// 初始化活动状态，所有查询都使用
	active := make([]bool, m)

	// 尝试逐个移除 queries，记录最大移除数量
	count := 0
	for i := 0; i < m; i++ {
		// 暂时移除第 i 个查询
		active[i] = true
		if check(active) {
			count++
		} else {
			// 如果不能转化，则恢复这个查询
			active[i] = false
		}
	}

	return count
}
