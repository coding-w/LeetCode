package main

func main() {

}

// NumMatrix https://leetcode.cn/problems/range-sum-query-2d-immutable/
type NumMatrix struct {
	nums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	tmp := make([][]int, len(matrix))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = make([]int, len(matrix[0])+1)
		sum := 0
		for j := 1; j < len(tmp[i]); j++ {
			sum += matrix[i][j-1]
			tmp[i][j] = sum
		}
	}
	return NumMatrix{
		tmp,
	}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.nums[i][col2+1] - this.nums[i][col1]
	}
	return sum
}
