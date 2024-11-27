package main

func main() {

}

// https://leetcode.cn/problems/find-the-maximum-number-of-fruits-collected/
func maxCollectedFruits(fruits [][]int) int {
	total := 0
	n := len(fruits)
	for i := 0; i < n; i++ {
		total += fruits[i][i]
	}
	total += fruits[0][n-1]
	for i := 1; i < n-1; i++ {
		total += max(fruits[i][n-1], fruits[i][n-2])
	}
	total += fruits[n-1][0]
	for i := 1; i < n-1; i++ {
		total += max(fruits[n-1][i], fruits[n-2][i])
	}
	return total + 2*fruits[n-1][n-1]
}
