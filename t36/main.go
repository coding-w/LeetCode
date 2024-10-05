package main

func main() {
	isValidSudoku([][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'9', '.', '.', '.', '3', '.', '.', '8', '.'},
		[]byte{'.', '.', '.', '8', '.', '3', '.', '.', '7'},
		[]byte{'.', '.', '.', '.', '4', '.', '.', '.', '.'},
		[]byte{'7', '.', '.', '.', '.', '.', '3', '.', '.'},
	})

}

// https://leetcode.cn/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	res := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		res[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			// 当前数字
			num := board[i][j] - '0' - 1
			// 当前在第几个 3*3宫格
			target := i/3*3 + j/3
			if row[i][num] || col[j][num] || res[target][num] {
				return false
			} else {
				col[j][num] = true
				row[i][num] = true
				res[i/3*3+j/3][num] = true
			}
		}
	}
	return true
}
