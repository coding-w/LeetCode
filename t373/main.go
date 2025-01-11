package main

import "fmt"

func main() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}

// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m := len(nums2)
	h := make(map[int]bool, m)
	for _, v := range nums2 {
		h[v] = true
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums1); i++ {
		target := k - nums1[i]
		if h[target] {
			res = append(res, []int{nums1[i], target})
		}
		target = nums1[i] - k
		if h[target] {
			res = append(res, []int{nums1[i], target})
		}
		target = nums1[i] + k
		if h[target] {
			res = append(res, []int{nums1[i], target})
		}
	}
	return res
}
