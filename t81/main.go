package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/description/
func search(nums []int, target int) bool {
	res := false
	for _, v := range nums {
		if v == target {
			res = true
			break
		}
	}
	return res
}
