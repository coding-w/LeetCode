package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

// https://leetcode.cn/problems/subsets/description/
func subsets(nums []int) [][]int {
	n := int(math.Pow(2, float64(len(nums))))
	res := make([][]int, n)
	j, k := 1, 1
	for _, num := range nums {
		for i := 0; i < k; i++ {
			tmp := append([]int{}, res[i]...)
			res[j] = append(tmp, num)
			j++
		}
		k = j
	}
	return res
}
