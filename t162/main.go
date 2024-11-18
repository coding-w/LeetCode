package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 4, 5}))
}

// https://leetcode.cn/problems/find-peak-element/
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r + l) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
