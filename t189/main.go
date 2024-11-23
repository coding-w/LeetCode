package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

// https://leetcode.cn/problems/rotate-array/
func rotate(nums []int, k int) {
	l := len(nums)
	k %= l
	rotated := append(nums[l-k:], nums[:l-k]...)
	copy(nums, rotated)
}
