package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

// https://leetcode.cn/problems/spiral-matrix/description/
// tips: 暴力破解，上下左右四个方向爬取
func spiralOrder(matrix [][]int) []int {
	// 结果数组
	res := make([]int, len(matrix)*len(matrix[0]))
	if len(matrix) == 1 {
		return matrix[0]
	}
	// 矩阵 长宽
	m, n := len(matrix), len(matrix[0])
	// 记录每行的开始位置
	x, y := 0, 0
	// 游标
	i, j := 0, 0
	// 已记录的条数
	k := 0
	// 右下左上
	d := 0
	for {
		switch d {
		case 0:
			{
				for ; j < n; j++ {
					res[k] = matrix[i][j]
					k++
				}
				if k == len(res) {
					return res
				}
				n -= 1
				j -= 1
				i = x + 1
				break
			}
		case 1:
			{
				for ; i < m; i++ {
					res[k] = matrix[i][j]
					k++
				}
				if k == len(res) {
					return res
				}
				m -= 1
				i -= 1
				j -= 1
				break
			}
		case 2:
			{
				for ; j >= y; j-- {
					res[k] = matrix[i][j]
					k++
				}
				if k == len(res) {
					return res
				}
				i -= 1
				j = y
				y += 1

				break
			}

		case 3:
			{
				for ; i > x; i-- {
					res[k] = matrix[i][j]
					k++
				}
				if k == len(res) {
					return res
				}
				x += 1
				i = x
				j = y
			}
		}
		d = (d + 1) % 4
	}

}
