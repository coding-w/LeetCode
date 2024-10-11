package main

import "fmt"

func main() {
	fmt.Println(totalNQueens(4))
}

// https://leetcode.cn/problems/n-queens-ii/description/
func totalNQueens(n int) int {
	var res int
	var dfs func(int, []int)
	dfs = func(i int, i2 []int) {
		if i == n {
			res++
		}
		for j := 0; j < n; j++ {
			if isValid(i, j, i2) {
				dfs(i+1, append(i2, j))
			}
		}
	}

	dfs(0, []int{})

	return res
}

func isValid(i, j int, arr []int) bool {

	for k := 0; k < i; k++ {
		if arr[k] == j || abs(arr[k]-j) == abs(k-i) {
			return false
		}
	}
	return true
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
