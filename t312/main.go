package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{8, 2, 6, 8, 9, 8, 1, 4, 1, 5, 3, 0, 7, 7, 0, 4, 2, 2, 5}))
}

// https://leetcode.cn/problems/burst-balloons/
// 超时....
func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	tmp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[i] = make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			tmp[i][j] = -1
		}
	}
	var recursion func(left, right int) int
	recursion = func(left, right int) int {
		if left > right-2 {
			return 0
		}
		if tmp[left][right] != -1 {
			return tmp[left][right]
		}
		for i := left + 1; i < right; i++ {
			sum := nums[left]*nums[i]*nums[right] + recursion(i, right) + recursion(left, i)
			tmp[left][right] = max(tmp[left][right], sum)
		}
		return tmp[left][right]
	}
	return recursion(0, len(nums)-1)
}
