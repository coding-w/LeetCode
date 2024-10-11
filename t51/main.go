package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

// https://leetcode.cn/problems/n-queens/description/
func solveNQueens(n int) [][]string {
	var res [][]string
	var dfs func(int, []int)
	dfs = func(i int, i2 []int) {
		if i == n {
			var tmp []string
			for j := 0; j < n; j++ {
				var s string
				for k := 0; k < n; k++ {
					if k == i2[j] {
						s += "Q"
					} else {
						s += "."
					}
				}
				tmp = append(tmp, s)
			}
			res = append(res, tmp)
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
