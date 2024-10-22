package main

import "fmt"

func main() {
	fmt.Println(combine(3, 2))
}

// https://leetcode.cn/problems/combinations/description/
func combine(n int, k int) [][]int {
	var result [][]int
	var path []int

	var dfs func(start int)
	dfs = func(start int) {
		// 如果当前组合的长度等于 k，添加到结果中
		if len(path) == k {
			// 复制 path 到 result 中
			combination := make([]int, k)
			copy(combination, path)
			result = append(result, combination)
			return
		}
		// 从 start 开始选择数字
		for i := start; i <= n; i++ {
			path = append(path, i)    // 选择 i
			dfs(i + 1)                // 递归
			path = path[:len(path)-1] // 撤销选择
		}
	}

	dfs(1) // 从 1 开始
	return result
}
