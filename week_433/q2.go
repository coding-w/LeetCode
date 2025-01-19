package main

import (
	"fmt"
	"sort"
)

// error
func minMaxSums(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	tmp := make([][]int, 0)
	tmp = append(tmp, []int{})
	// 外层循环遍历所有可能的起始点
	for i := 0; i < len(nums); i++ {
		t := len(tmp)
		for j := 0; j < t; j++ {
			if len(tmp[j]) < k {
				tmp[j] = append([]int{}, tmp[j]...)
				tmp = append(tmp, append(tmp[j], nums[i]))
			}
		}

	}
	tmp = tmp[1:]
	for i := 0; i < len(tmp); i++ {
		m := len(tmp[i])
		res += tmp[i][0] + tmp[i][m-1]

	}
	return res
}

func main() {
	arr := []int{2, 0, 1, 3, 2}
	fmt.Println(minMaxSums(arr, 5))
}
