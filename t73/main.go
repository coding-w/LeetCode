package main

func main() {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
}

// https://leetcode.cn/problems/set-matrix-zeroes/description/
// 先统计一次 0 的i,j再构造结果
func setZeroes(matrix [][]int) {
	iMap := make(map[int]bool)
	jMap := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				iMap[i] = true
				jMap[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, oki := iMap[i]
			_, okj := jMap[j]
			if oki || okj {
				matrix[i][j] = 0
				continue
			}
		}
	}
	//fmt.Println(matrix)
}
