package main

func main() {
	//r := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	r := [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}
	solve(r)
}

// https://leetcode.cn/problems/surrounded-regions/
func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	if m < 3 || n < 3 {
		return
	}
	tmp := make([][]bool, m+2)
	for i := 0; i < m+2; i++ {
		tmp[i] = make([]bool, n+2)
	}
	var dfs func(i, j int)
	d := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	dfs = func(i, j int) {
		if board[i][j] != 'O' && !tmp[i+1][j+1] {
			return
		}
		for _, item := range d {
			i1 := i + item[0]
			j1 := j + item[1]
			if i1 > 0 && j1 > 0 && i1 < m && j1 < n && board[i1][j1] == 'O' && tmp[i+1][j+1] {
				tmp[i1+1][j1+1] = true
				board[i][j] = 'V'
				dfs(i1, j1)
			}
		}
	}
	// 遍历外围
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' && !tmp[i+1][j+1] {
				tmp[i+1][j+1] = true
				dfs(i, j)
			}
		}

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !tmp[i+1][j+1] {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
	}
	//fmt.Println(board)
}
