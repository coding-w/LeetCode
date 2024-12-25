package main

import "fmt"

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
}

// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/
// dfs 加 备忘录
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(i, j, val int) int
	dfs = func(i, j, val int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if matrix[i][j] <= val {
			return 0
		}
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		val = matrix[i][j]
		left := dfs(i-1, j, val)
		right := dfs(i+1, j, val)
		top := dfs(i, j-1, val)
		bottom := dfs(i, j+1, val)
		memo[i][j] = max(left, right, top, bottom) + 1
		return memo[i][j]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			left := dfs(i-1, j, matrix[i][j])
			right := dfs(i+1, j, matrix[i][j])
			top := dfs(i, j-1, matrix[i][j])
			bottom := dfs(i, j+1, matrix[i][j])
			memo[i][j] = max(left, right, top, bottom) + 1
		}
	}
	res := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, memo[i][j])
		}
	}
	return res
}
