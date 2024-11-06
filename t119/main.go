package main

import "fmt"

func main() {
	fmt.Println(getRow(0))
}

// https://leetcode.cn/problems/pascals-triangle-ii/description/
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res[rowIndex]
}
