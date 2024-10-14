package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

// https://leetcode.cn/problems/spiral-matrix-ii/description/
// tips: 暴力破解，上下左右四个方向爬取
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	// 结果数组
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// 矩阵 长宽
	m, n1 := n, n
	// 记录每行的开始位置
	x, y := 0, 0
	// 游标
	i, j := 0, 0
	// 已记录的条数
	k := 1
	// 右下左上
	d := 0
	for {
		switch d {
		case 0:
			{
				for ; j < n1; j++ {
					res[i][j] = k
					k++
				}
				if k == n*n+1 {
					return res
				}
				n1 -= 1
				j -= 1
				i = x + 1
				break
			}
		case 1:
			{
				for ; i < m; i++ {
					res[i][j] = k
					k++
				}
				if k == n*n+1 {
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
					res[i][j] = k
					k++
				}
				if k == n*n+1 {
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
					res[i][j] = k
					k++
				}
				if k == n*n+1 {
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
