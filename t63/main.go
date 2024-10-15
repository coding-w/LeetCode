package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}

// https://leetcode.cn/problems/unique-paths/description/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 将有障碍物的标识出来
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}
	// 初始化迷宫
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == -1 {
			break
		}
		obstacleGrid[i][0] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == -1 {
			break
		}
		obstacleGrid[0][i] = 1
	}
	// 走迷宫
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 遇到障碍物跳过
			if obstacleGrid[i][j] == -1 {
				continue
			}
			obstacleGrid[i][j] = max(obstacleGrid[i-1][j], 0) + max(obstacleGrid[i][j-1], 0)
		}
	}
	// 障碍物在终点
	if obstacleGrid[m-1][n-1] == -1 {
		return 0
	}
	// 返回结果
	return obstacleGrid[m-1][n-1]
}
