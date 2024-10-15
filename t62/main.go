package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(19, 13))
}

// https://leetcode.cn/problems/unique-paths/description/
// 动态规划
func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		res[i][0] = 1
	}
	for i := 0; i < n; i++ {
		res[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}

// 会超时
func uniquePaths1(m int, n int) int {
	var dfs func(int, int)
	res := 0
	dfs = func(i int, j int) {
		if i == m && j == n {
			res++
			return
		}
		if i < m {
			dfs(i+1, j)
		}
		if j < n {
			dfs(i, j+1)
		}
	}
	dfs(1, 1)
	return res
}
