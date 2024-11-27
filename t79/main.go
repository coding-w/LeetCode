package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCB"))
}

// https://leetcode.cn/problems/word-search/
func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	isRun := make([][]bool, n)
	for i := range isRun {
		isRun[i] = make([]bool, m)
	}
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		if index == len(word) {
			return true
		}
		isRun[i][j] = true
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if (x < 0 || x > n-1 || y < 0 || y > m-1) || isRun[x][y] {
				continue
			}
			if board[x][y] == word[index] {
				if dfs(x, y, index+1) {
					return true
				}
				isRun[x][y] = false
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 1) {
					return true
				}
				isRun[i][j] = false
			}
		}
	}
	return false
}
