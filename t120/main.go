package main

func main() {

}

// https://leetcode.cn/problems/triangle/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	for i := len(triangle) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			triangle[i-1][j] += min(triangle[i][j], triangle[i][j+1])
		}
	}
	return triangle[0][0]
}
