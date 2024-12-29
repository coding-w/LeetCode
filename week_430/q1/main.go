package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([][]int{{3, 2}, {1, 3}, {3, 4}, {0, 1}}))
}

func minimumOperations(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			k := grid[i][j] - grid[i-1][j]
			if k <= 0 {
				grid[i][j] += -k + 1
				res += -k + 1
			}
		}
	}
	return res
}
