package main

import (
	"sort"
)

func main() {
	merge([]int{0}, 0, []int{1}, 1)
}

// https://leetcode.cn/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for j := 0; j < n; j++ {
		nums1[m+j] = nums2[j]
	}
	sort.Ints(nums1)
	//fmt.Println(nums1)
}
