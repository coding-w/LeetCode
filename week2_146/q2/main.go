package main

import "fmt"

const MOD = 1e9 + 7

func countPathsWithXorValue(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	// dp[i][j][xor_val] 表示到达位置 (i, j) 时，XOR 值为 xor_val 的路径数
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 16) // XOR 的值在 [0, 16] 范围内
		}
	}

	// 初始状态，起点 (0, 0)
	dp[0][0][grid[0][0]] = 1

	// 遍历每个格子
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for xorVal := 0; xorVal < 16; xorVal++ {
				if dp[i][j][xorVal] > 0 {
					// 向下走
					if i+1 < m {
						dp[i+1][j][xorVal^grid[i+1][j]] = (dp[i+1][j][xorVal^grid[i+1][j]] + dp[i][j][xorVal]) % MOD
					}
					// 向右走
					if j+1 < n {
						dp[i][j+1][xorVal^grid[i][j+1]] = (dp[i][j+1][xorVal^grid[i][j+1]] + dp[i][j][xorVal]) % MOD
					}
				}
			}
		}
	}

	// 返回终点 (m-1, n-1) 的 XOR 值为 k 的路径数目
	return dp[m-1][n-1][k]
}

func main() {
	// 示例
	grid := [][]int{
		{2, 1, 5},
		{7, 10, 0},
		{12, 6, 4},
	}
	k := 11
	fmt.Println(countPathsWithXorValue(grid, k)) // 输出路径数目
}
