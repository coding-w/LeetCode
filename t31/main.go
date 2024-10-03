package main

import (
	"sort"
)

func main() {
	nextPermutation([]int{1, 2, 3})
}

// https://leetcode.cn/problems/next-permutation/description/
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	n := len(nums)
	l, r := 0, n-1
	for i := 0; i < n-1; i++ {
		// 找到第一个小于右边的数
		if nums[i] < nums[i+1] {
			l = i
		}
	}

	for j := l + 1; j < n; j++ {
		// 找到第一个大于左边的数
		if nums[j] > nums[l] {
			r = j
		}
	}
	// 交换
	nums[l], nums[r] = nums[r], nums[l]
	ss := nums[l+1:]
	sort.Ints(ss)
	for i := 0; i < len(ss); i++ {
		nums[l+1+i] = ss[i]
	}

	for _, v := range nums {
		print(v, " ")
	}

}
