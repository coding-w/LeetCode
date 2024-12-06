package main

func main() {

}

// https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix) - 1
	j := 0
	for n >= 0 && j < len(matrix[0]) {
		if matrix[n][j] == target {
			return true
		} else if matrix[n][j] > target {
			n--
		} else {
			j++
		}
	}
	return false
}
