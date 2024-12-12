package main

func main() {

}

// https://leetcode.cn/problems/game-of-life/
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := 0
			for k := 0; k < 8; k++ {
				x, y := i+dir[k][0], j+dir[k][1]
				if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 1 {
					cnt++
				}
			}
			if board[i][j] == 1 && (cnt == 2 || cnt == 3) {
				res[i][j] = 1
			}
			if board[i][j] == 0 && cnt == 3 {
				res[i][j] = 1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = res[i][j]
		}
	}
}
