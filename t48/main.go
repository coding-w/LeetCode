package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}

// https://leetcode.cn/problems/rotate-image/description/
// 顺时针旋转90°：matrix[i][j] -> matrix[j][n−1−i]
// = 先上下对称翻转：matrix[i][j] -> matrix[n-i-1][j]
//   - 再主对角线翻转：matrix[n-i-1][j] -> matrix[j][n−1−i]
func rotate(matrix [][]int) {
	m := len(matrix)
	for i := 0; i < m/2; i++ {
		matrix[i], matrix[m-1-i] = matrix[m-1-i], matrix[i]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
