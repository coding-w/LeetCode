package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))

}

// https://leetcode.cn/problems/minimum-path-sum/description/
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	res[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		res[i][0] = res[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		res[0][j] = res[0][j-1] + grid[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			res[i][j] = min(res[i][j-1], res[i-1][j]) + grid[i][j]
		}
	}
	return res[n-1][m-1]
}
