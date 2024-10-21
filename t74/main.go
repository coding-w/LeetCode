package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))

}

// https://leetcode.cn/problems/search-a-2d-matrix/description/
// tips：二分法
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	left, right := 0, m*n
	for left < right {
		mid := (right + left) / 2
		if matrix[mid/n][mid%n] == target {
			return true
		}
		if matrix[mid/n][mid%n] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return matrix[left/n][left%n] == target
}
